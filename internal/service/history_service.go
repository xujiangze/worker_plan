package service

import (
	"fmt"
	"worker_plan/internal/model"
	"worker_plan/internal/repository"
)

// HistoryService 历史服务
type HistoryService struct {
	planHistoryRepo repository.PlanHistoryRepository
}

// NewHistoryService 创建历史服务
func NewHistoryService(planHistoryRepo repository.PlanHistoryRepository) *HistoryService {
	return &HistoryService{
		planHistoryRepo: planHistoryRepo,
	}
}

// GetHistoryByPlanID 获取计划的历史记录
func (s *HistoryService) GetHistoryByPlanID(planID uint, page, pageSize int) ([]*model.PlanHistory, int64, error) {
	offset := (page - 1) * pageSize

	histories, total, err := s.planHistoryRepo.FindByPlanID(planID, offset, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get history: %w", err)
	}

	return histories, total, nil
}

// RecordHistory 记录变更历史
func (s *HistoryService) RecordHistory(planID uint, fieldName, oldValue, newValue, changeType string) error {
	history := &model.PlanHistory{
		PlanID:     planID,
		FieldName:  fieldName,
		OldValue:   oldValue,
		NewValue:   newValue,
		ChangeType: changeType,
	}

	if err := s.planHistoryRepo.Create(history); err != nil {
		return fmt.Errorf("failed to record history: %w", err)
	}

	return nil
}
