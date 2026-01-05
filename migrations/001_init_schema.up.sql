-- 创建 plans 表
CREATE TABLE IF NOT EXISTS plans (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    priority VARCHAR(20) NOT NULL CHECK (priority IN ('High', 'Medium', 'Low')),
    status VARCHAR(20) NOT NULL CHECK (status IN ('Todo', 'InProgress', 'Done', 'Cancelled')),
    due_date TIMESTAMP,
    progress INTEGER NOT NULL DEFAULT 0 CHECK (progress >= 0 AND progress <= 100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_plans_status ON plans(status);
CREATE INDEX IF NOT EXISTS idx_plans_priority ON plans(priority);
CREATE INDEX IF NOT EXISTS idx_plans_due_date ON plans(due_date);
CREATE INDEX IF NOT EXISTS idx_plans_deleted_at ON plans(deleted_at);

-- 创建 plan_histories 表
CREATE TABLE IF NOT EXISTS plan_histories (
    id SERIAL PRIMARY KEY,
    plan_id INTEGER NOT NULL,
    field_name VARCHAR(50) NOT NULL,
    old_value TEXT,
    new_value TEXT,
    change_type VARCHAR(20) NOT NULL CHECK (change_type IN ('Status', 'Progress', 'Info')),
    changed_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_plan_histories_plan_id FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE CASCADE
);

-- 创建索引
CREATE INDEX IF NOT EXISTS idx_plan_histories_plan_id ON plan_histories(plan_id);
CREATE INDEX IF NOT EXISTS idx_plan_histories_changed_at ON plan_histories(changed_at);
CREATE INDEX IF NOT EXISTS idx_plan_histories_change_type ON plan_histories(change_type);

-- 创建更新时间触发器函数
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- 创建触发器
CREATE TRIGGER update_plans_updated_at BEFORE UPDATE ON plans
    FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
