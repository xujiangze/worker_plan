# 部署文档

## 环境要求

- Go 1.21 或更高版本
- PostgreSQL 12 或更高版本
- 操作系统: Linux, macOS, Windows

## 环境配置

### 1. 配置环境变量

复制 `.env.example` 文件为 `.env`:

```bash
cp .env.example .env
```

编辑 `.env` 文件,配置以下参数:

```env
# 服务器配置
SERVER_PORT=8080
SERVER_MODE=debug

# 数据库配置
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=worker_plan
DB_SSLMODE=disable

# 日志配置
LOG_LEVEL=info
LOG_FORMAT=console
```

### 2. 数据库配置

#### 创建数据库

```sql
-- 连接到 PostgreSQL
psql -U postgres

-- 创建数据库
CREATE DATABASE worker_plan;

-- 创建用户(可选)
CREATE USER worker_plan_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE worker_plan TO worker_plan_user;
```

#### 数据库迁移

项目使用 GORM 的 AutoMigrate 功能自动创建表结构。首次启动时会自动执行迁移。

如果需要手动执行 SQL 迁移脚本:

```bash
# 执行迁移
psql -U postgres -d worker_plan -f migrations/001_init_schema.up.sql

# 回滚迁移
psql -U postgres -d worker_plan -f migrations/001_init_schema.down.sql
```

## 本地开发

### 1. 安装依赖

```bash
go mod download
```

### 2. 运行服务

```bash
go run cmd/server/main.go
```

服务将在 `http://localhost:8080` 启动。

### 3. 测试 API

```bash
# 健康检查
curl http://localhost:8080/health

# 创建计划
curl -X POST http://localhost:8080/api/plans \
  -H "Content-Type: application/json" \
  -d '{
    "title": "完成项目文档",
    "description": "编写项目的技术文档和用户手册",
    "priority": "High"
  }'
```

## 构建和部署

### 1. 构建可执行文件

```bash
# 构建当前平台
go build -o worker_plan cmd/server/main.go

# 构建 Linux 平台
GOOS=linux GOARCH=amd64 go build -o worker_plan-linux cmd/server/main.go

# 构建 macOS 平台
GOOS=darwin GOARCH=amd64 go build -o worker_plan-macos cmd/server/main.go

# 构建 Windows 平台
GOOS=windows GOARCH=amd64 go build -o worker_plan.exe cmd/server/main.go
```

### 2. 部署到服务器

#### 使用 systemd 部署(Linux)

创建服务文件 `/etc/systemd/system/worker-plan.service`:

```ini
[Unit]
Description=Worker Plan Management System
After=network.target postgresql.service

[Service]
Type=simple
User=worker_plan
WorkingDirectory=/opt/worker_plan
ExecStart=/opt/worker_plan/worker_plan
Restart=on-failure
RestartSec=5s
Environment="SERVER_PORT=8080"
Environment="DB_HOST=localhost"
Environment="DB_PORT=5432"
Environment="DB_USER=worker_plan_user"
Environment="DB_PASSWORD=your_password"
Environment="DB_NAME=worker_plan"
Environment="DB_SSLMODE=disable"
Environment="LOG_LEVEL=info"
Environment="LOG_FORMAT=json"

[Install]
WantedBy=multi-user.target
```

启动服务:

```bash
# 重载 systemd 配置
sudo systemctl daemon-reload

# 启动服务
sudo systemctl start worker-plan

# 设置开机自启
sudo systemctl enable worker-plan

# 查看服务状态
sudo systemctl status worker-plan

# 查看服务日志
sudo journalctl -u worker-plan -f
```

#### 使用 Docker 部署

创建 `Dockerfile`:

```dockerfile
# 构建阶段
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o worker_plan cmd/server/main.go

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/worker_plan .

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./worker_plan"]
```

构建和运行 Docker 容器:

```bash
# 构建镜像
docker build -t worker-plan:latest .

# 运行容器
docker run -d \
  --name worker-plan \
  -p 8080:8080 \
  -e SERVER_PORT=8080 \
  -e DB_HOST=host.docker.internal \
  -e DB_PORT=5432 \
  -e DB_USER=postgres \
  -e DB_PASSWORD=your_password \
  -e DB_NAME=worker_plan \
  -e DB_SSLMODE=disable \
  -e LOG_LEVEL=info \
  worker-plan:latest

# 查看日志
docker logs -f worker-plan

# 停止容器
docker stop worker-plan

# 删除容器
docker rm worker-plan
```

#### 使用 Docker Compose 部署

创建 `docker-compose.yml`:

```yaml
version: '3.8'

services:
  postgres:
    image: postgres:15-alpine
    container_name: worker-plan-db
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: your_password
      POSTGRES_DB: worker_plan
    ports:
      - "5432:5432"
    volumes:
      - postgres-data:/var/lib/postgresql/data
    networks:
      - worker-plan-network

  app:
    build: .
    container_name: worker-plan-app
    ports:
      - "8080:8080"
    environment:
      SERVER_PORT: 8080
      DB_HOST: postgres
      DB_PORT: 5432
      DB_USER: postgres
      DB_PASSWORD: your_password
      DB_NAME: worker_plan
      DB_SSLMODE: disable
      LOG_LEVEL: info
    depends_on:
      - postgres
    networks:
      - worker-plan-network

volumes:
  postgres-data:

networks:
  worker-plan-network:
    driver: bridge
```

启动服务:

```bash
# 启动所有服务
docker-compose up -d

# 查看日志
docker-compose logs -f

# 停止所有服务
docker-compose down

# 停止并删除数据卷
docker-compose down -v
```

## 反向代理配置

### Nginx 配置

创建 Nginx 配置文件 `/etc/nginx/sites-available/worker-plan`:

```nginx
upstream worker_plan {
    server localhost:8080;
}

server {
    listen 80;
    server_name your-domain.com;

    location / {
        proxy_pass http://worker_plan;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

启用配置:

```bash
# 创建符号链接
sudo ln -s /etc/nginx/sites-available/worker-plan /etc/nginx/sites-enabled/

# 测试配置
sudo nginx -t

# 重载 Nginx
sudo systemctl reload nginx
```

### HTTPS 配置(使用 Let's Encrypt)

```bash
# 安装 certbot
sudo apt-get install certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d your-domain.com

# 自动续期
sudo certbot renew --dry-run
```

## 监控和日志

### 日志管理

日志输出到标准输出,可以使用以下工具收集:

- **journalctl**: systemd 日志
- **docker logs**: Docker 容器日志
- **ELK Stack**: Elasticsearch, Logstash, Kibana
- **Loki**: Grafana Loki

### 健康检查

定期检查服务健康状态:

```bash
curl http://localhost:8080/health
```

### 性能监控

可以使用以下工具监控性能:

- **Prometheus + Grafana**: 指标收集和可视化
- **New Relic**: APM 监控
- **Datadog**: 基础设施监控

## 备份和恢复

### 数据库备份

```bash
# 备份数据库
pg_dump -U postgres -d worker_plan > backup_$(date +%Y%m%d_%H%M%S).sql

# 恢复数据库
psql -U postgres -d worker_plan < backup_20240101_120000.sql
```

### 自动备份脚本

创建备份脚本 `/opt/backup.sh`:

```bash
#!/bin/bash

BACKUP_DIR="/opt/backups"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/worker_plan_$DATE.sql"

mkdir -p $BACKUP_DIR

pg_dump -U postgres -d worker_plan > $BACKUP_FILE

# 保留最近 7 天的备份
find $BACKUP_DIR -name "worker_plan_*.sql" -mtime +7 -delete

echo "Backup completed: $BACKUP_FILE"
```

添加到 crontab:

```bash
# 每天凌晨 2 点备份
0 2 * * * /opt/backup.sh >> /var/log/backup.log 2>&1
```

## 故障排查

### 常见问题

#### 1. 数据库连接失败

检查数据库配置和网络连接:

```bash
# 测试数据库连接
psql -U postgres -h localhost -p 5432 -d worker_plan

# 检查 PostgreSQL 服务状态
sudo systemctl status postgresql
```

#### 2. 端口被占用

检查端口占用:

```bash
# Linux/macOS
lsof -i :8080

# 杀死占用端口的进程
kill -9 <PID>
```

#### 3. 权限问题

确保应用有正确的文件权限:

```bash
# 修改文件所有者
sudo chown -R worker_plan:worker_plan /opt/worker_plan

# 修改文件权限
sudo chmod -R 755 /opt/worker_plan
```

## 安全建议

1. **使用强密码**: 数据库密码应使用强密码
2. **启用 SSL**: 生产环境应启用数据库 SSL 连接
3. **防火墙配置**: 只开放必要的端口
4. **定期更新**: 定期更新系统和依赖包
5. **日志审计**: 定期检查日志,发现异常行为
6. **备份策略**: 制定并执行定期备份策略
