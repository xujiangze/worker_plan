-- 删除触发器
DROP TRIGGER IF EXISTS update_plans_updated_at ON plans;

-- 删除触发器函数
DROP FUNCTION IF EXISTS update_updated_at_column();

-- 删除 plan_histories 表
DROP TABLE IF EXISTS plan_histories;

-- 删除 plans 表
DROP TABLE IF EXISTS plans;
