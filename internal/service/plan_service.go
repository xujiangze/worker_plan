package service

import (
	"errors"
	"fmt"
	"worker_plan/internal/model"
	"worker_plan/internal/repository"
)

// PlanService 计划服务
type PlanService struct {
	planRepo        repository.PlanRepository
	planHistoryRepo repository.PlanHistoryRepository
}

// NewPlanService 创建计划服务
func NewPlanService(planRepo repository.PlanRepository, planHistoryRepo repository.PlanHistoryRepository) *PlanService {
	return &PlanService{
		planRepo:        planRepo,
		planHistoryRepo: planHistoryRepo,
	}
}

// CreatePlanRequest 创建计划请求
type CreatePlanRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description"`
	Priority    string  `json:"priority" binding:"required,oneof=High Medium Low"`
	DueDate     *string `json:"due_date"`
}

// UpdatePlanRequest 更新计划请求
type UpdatePlanRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Priority    *string `json:"priority" binding:"omitempty,oneof=High Medium Low"`
	DueDate     *string `json:"due_date"`
}

// CreatePlan 创建计划
func (s *PlanService) CreatePlan(req *CreatePlanRequest) (*model.Plan, error) {
	// 创建计划模型
	plan := &model.Plan{
		Title:       req.Title,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      "Todo",
		Progress:    0,
	}

	// 保存到数据库
	if err := s.planRepo.Create(plan); err != nil {
		return nil, fmt.Errorf("failed to create plan: %w", err)
	}

	return plan, nil
}

// GetPlan 获取计划
func (s *PlanService) GetPlan(id uint) (*model.Plan, error) {
	plan, err := s.planRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, fmt.Errorf("failed to get plan: %w", err)
	}
	return plan, nil
}

// GetPlans 获取计划列表
func (s *PlanService) GetPlans(page, pageSize int, status, priority string) ([]*model.Plan, int64, error) {
	offset := (page - 1) * pageSize

	// 构建筛选条件
	filters := make(map[string]interface{})
	if status != "" {
		filters["status = ?"] = status
	}
	if priority != "" {
		filters["priority = ?"] = priority
	}

	plans, total, err := s.planRepo.FindAll(offset, pageSize, filters, "created_at DESC")
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get plans: %w", err)
	}

	return plans, total, nil
}

// UpdatePlan 更新计划
func (s *PlanService) UpdatePlan(id uint, req *UpdatePlanRequest) (*model.Plan, error) {
	// 获取现有计划
	plan, err := s.planRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, fmt.Errorf("failed to get plan: %w", err)
	}

	// 记录变更历史
	if req.Title != nil && *req.Title != plan.Title {
		s.recordHistory(plan.ID, "title", plan.Title, *req.Title, model.ChangeTypeInfo)
		plan.Title = *req.Title
	}

	if req.Description != nil && *req.Description != plan.Description {
		s.recordHistory(plan.ID, "description", plan.Description, *req.Description, model.ChangeTypeInfo)
		plan.Description = *req.Description
	}

	if req.Priority != nil && *req.Priority != plan.Priority {
		s.recordHistory(plan.ID, "priority", plan.Priority, *req.Priority, model.ChangeTypeInfo)
		plan.Priority = *req.Priority
	}

	if req.DueDate != nil {
		// 这里需要处理时间格式转换
		// 简化处理,实际应该解析时间字符串
		// plan.DueDate = parseTime(*req.DueDate)
	}

	// 更新计划
	if err := s.planRepo.Update(plan); err != nil {
		return nil, fmt.Errorf("failed to update plan: %w", err)
	}

	return plan, nil
}

// DeletePlan 删除计划
func (s *PlanService) DeletePlan(id uint) error {
	// 检查计划是否存在
	_, err := s.planRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return errors.New("plan not found")
		}
		return fmt.Errorf("failed to get plan: %w", err)
	}

	// 软删除计划
	if err := s.planRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete plan: %w", err)
	}

	return nil
}

// recordHistory 记录变更历史
func (s *PlanService) recordHistory(planID uint, fieldName, oldValue, newValue, changeType string) {
	history := &model.PlanHistory{
		PlanID:     planID,
		FieldName:  fieldName,
		OldValue:   oldValue,
		NewValue:   newValue,
		ChangeType: changeType,
	}
	s.planHistoryRepo.Create(history)
}
