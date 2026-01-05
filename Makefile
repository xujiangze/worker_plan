.PHONY: help install-backend install-frontend install dev-backend dev-frontend dev build-backend build-frontend build clean

# 默认目标
help:
	@echo "可用的命令:"
	@echo "  make install          - 安装所有依赖（后端和前端）"
	@echo "  make install-backend   - 安装后端依赖"
	@echo "  make install-frontend  - 安装前端依赖"
	@echo "  make dev               - 同时启动后端和前端开发服务器"
	@echo "  make dev-backend       - 启动后端开发服务器"
	@echo "  make dev-frontend      - 启动前端开发服务器"
	@echo "  make build             - 构建所有项目（后端和前端）"
	@echo "  make build-backend     - 构建后端"
	@echo "  make build-frontend    - 构建前端"
	@echo "  make clean             - 清理构建文件"
	@echo "  make run               - 运行已构建的后端"

# 安装所有依赖
install: install-backend install-frontend
	@echo "✓ 所有依赖安装完成"

# 安装后端依赖
install-backend:
	@echo "安装后端依赖..."
	@go mod download
	@echo "✓ 后端依赖安装完成"

# 安装前端依赖
install-frontend:
	@echo "安装前端依赖..."
	@cd frontend && npm install
	@echo "✓ 前端依赖安装完成"

# 同时启动后端和前端开发服务器
dev:
	@echo "启动后端和前端开发服务器..."
	@make -j2 dev-backend dev-frontend

# 启动后端开发服务器
dev-backend:
	@echo "启动后端开发服务器 (端口 8080)..."
	@if [ ! -f .env ]; then \
		echo "警告: .env 文件不存在，从 .env.example 复制..."; \
		cp .env.example .env; \
	fi
	@go run cmd/server/main.go

# 启动前端开发服务器
dev-frontend:
	@echo "启动前端开发服务器 (端口 5173)..."
	@cd frontend && npm run dev

# 构建所有项目
build: build-backend build-frontend
	@echo "✓ 所有项目构建完成"

# 构建后端
build-backend:
	@echo "构建后端..."
	@go build -o bin/server cmd/server/main.go
	@echo "✓ 后端构建完成: bin/server"

# 构建前端
build-frontend:
	@echo "构建前端..."
	@cd frontend && npm run build
	@echo "✓ 前端构建完成: frontend/dist"

# 运行已构建的后端
run:
	@echo "运行后端服务器..."
	@if [ ! -f .env ]; then \
		echo "警告: .env 文件不存在，从 .env.example 复制..."; \
		cp .env.example .env; \
	fi
	@./bin/server

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf bin/
	@rm -rf frontend/dist/
	@rm -rf frontend/node_modules/.vite
	@echo "✓ 清理完成"

# 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...
	@cd frontend && npm run format 2>/dev/null || echo "前端格式化命令未配置"
	@echo "✓ 代码格式化完成"

# 运行测试
test:
	@echo "运行测试..."
	@go test -v ./...
	@echo "✓ 测试完成"

# 检查代码
lint:
	@echo "检查代码..."
	@go vet ./...
	@cd frontend && npm run lint 2>/dev/null || echo "前端 lint 命令未配置"
	@echo "✓ 代码检查完成"
