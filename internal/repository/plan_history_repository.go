package repository

import (
	"worker_plan/internal/model"

	"gorm.io/gorm"
)

// planHistoryRepository 计划历史仓储实现
type planHistoryRepository struct {
	db *gorm.DB
}

// NewPlanHistoryRepository 创建计划历史仓储
func NewPlanHistoryRepository(db *gorm.DB) PlanHistoryRepository {
	return &planHistoryRepository{db: db}
}

// Create 创建历史记录
func (r *planHistoryRepository) Create(history *model.PlanHistory) error {
	return r.db.Create(history).Error
}

// FindByPlanID 根据计划 ID 查找历史记录(支持分页)
func (r *planHistoryRepository) FindByPlanID(planID uint, offset, limit int) ([]*model.PlanHistory, int64, error) {
	var histories []*model.PlanHistory
	var total int64

	query := r.db.Model(&model.PlanHistory{}).Where("plan_id = ?", planID)

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用分页和排序
	if err := query.Order("changed_at DESC").Offset(offset).Limit(limit).Find(&histories).Error; err != nil {
		return nil, 0, err
	}

	return histories, total, nil
}

// CountByPlanID 统计计划的历史记录数量
func (r *planHistoryRepository) CountByPlanID(planID uint) (int64, error) {
	var count int64
	if err := r.db.Model(&model.PlanHistory{}).Where("plan_id = ?", planID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
