package service

import (
	"errors"
	"fmt"
	"worker_plan/internal/model"
	"worker_plan/internal/repository"
)

// ProgressService 进度服务
type ProgressService struct {
	planRepo        repository.PlanRepository
	planHistoryRepo repository.PlanHistoryRepository
}

// NewProgressService 创建进度服务
func NewProgressService(planRepo repository.PlanRepository, planHistoryRepo repository.PlanHistoryRepository) *ProgressService {
	return &ProgressService{
		planRepo:        planRepo,
		planHistoryRepo: planHistoryRepo,
	}
}

// UpdateStatusRequest 更新状态请求
type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=Todo InProgress Done Cancelled"`
}

// UpdateProgressRequest 更新进度请求
type UpdateProgressRequest struct {
	Progress int `json:"progress" binding:"required,min=0,max=100"`
}

// validStatusTransitions 有效的状态转换
var validStatusTransitions = map[string][]string{
	"Todo":       {"InProgress", "Cancelled"},
	"InProgress": {"Done", "Cancelled", "Todo"},
	"Done":       {"InProgress", "Todo"},
	"Cancelled":  {"Todo", "InProgress"},
}

// UpdateStatus 更新状态
func (s *ProgressService) UpdateStatus(id uint, req *UpdateStatusRequest) (*model.Plan, error) {
	// 获取现有计划
	plan, err := s.planRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, fmt.Errorf("failed to get plan: %w", err)
	}

	// 验证状态转换
	if !isValidStatusTransition(plan.Status, req.Status) {
		return nil, fmt.Errorf("invalid status transition from %s to %s", plan.Status, req.Status)
	}

	// 记录历史
	oldStatus := plan.Status
	plan.Status = req.Status

	// 如果状态变为 Done,自动设置进度为 100%
	if req.Status == "Done" {
		plan.Progress = 100
	}

	// 更新计划
	if err := s.planRepo.Update(plan); err != nil {
		return nil, fmt.Errorf("failed to update plan: %w", err)
	}

	// 记录状态变更历史
	s.recordHistory(plan.ID, "status", oldStatus, req.Status, model.ChangeTypeStatus)

	return plan, nil
}

// UpdateProgress 更新进度
func (s *ProgressService) UpdateProgress(id uint, req *UpdateProgressRequest) (*model.Plan, error) {
	// 获取现有计划
	plan, err := s.planRepo.FindByID(id)
	if err != nil {
		if errors.Is(err, repository.ErrRecordNotFound) {
			return nil, errors.New("plan not found")
		}
		return nil, fmt.Errorf("failed to get plan: %w", err)
	}

	// 记录历史
	oldProgress := plan.Progress
	plan.Progress = req.Progress

	// 如果进度达到 100%,自动更新状态为 Done
	if req.Progress == 100 && plan.Status != "Done" {
		oldStatus := plan.Status
		plan.Status = "Done"
		s.recordHistory(plan.ID, "status", oldStatus, "Done", model.ChangeTypeStatus)
	}

	// 如果进度小于 100% 且状态为 Done,自动更新状态为 InProgress
	if req.Progress < 100 && plan.Status == "Done" {
		oldStatus := plan.Status
		plan.Status = "InProgress"
		s.recordHistory(plan.ID, "status", oldStatus, "InProgress", model.ChangeTypeStatus)
	}

	// 更新计划
	if err := s.planRepo.Update(plan); err != nil {
		return nil, fmt.Errorf("failed to update plan: %w", err)
	}

	// 记录进度变更历史
	s.recordHistory(plan.ID, "progress", fmt.Sprintf("%d", oldProgress), fmt.Sprintf("%d", req.Progress), model.ChangeTypeProgress)

	return plan, nil
}

// isValidStatusTransition 验证状态转换是否有效
func isValidStatusTransition(from, to string) bool {
	if from == to {
		return true
	}

	validTransitions, exists := validStatusTransitions[from]
	if !exists {
		return false
	}

	for _, validTo := range validTransitions {
		if validTo == to {
			return true
		}
	}

	return false
}

// recordHistory 记录变更历史
func (s *ProgressService) recordHistory(planID uint, fieldName, oldValue, newValue, changeType string) {
	history := &model.PlanHistory{
		PlanID:     planID,
		FieldName:  fieldName,
		OldValue:   oldValue,
		NewValue:   newValue,
		ChangeType: changeType,
	}
	s.planHistoryRepo.Create(history)
}
