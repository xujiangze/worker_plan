package repository

import (
	"errors"
	"worker_plan/internal/model"

	"gorm.io/gorm"
)

// planRepository 计划仓储实现
type planRepository struct {
	db *gorm.DB
}

// NewPlanRepository 创建计划仓储
func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db: db}
}

// Create 创建计划
func (r *planRepository) Create(plan *model.Plan) error {
	return r.db.Create(plan).Error
}

// FindByID 根据 ID 查找计划
func (r *planRepository) FindByID(id uint) (*model.Plan, error) {
	var plan model.Plan
	err := r.db.First(&plan, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, err
	}
	return &plan, nil
}

// FindAll 查找所有计划(支持筛选、排序、分页)
func (r *planRepository) FindAll(offset, limit int, filters map[string]interface{}, orderBy string) ([]*model.Plan, int64, error) {
	var plans []*model.Plan
	var total int64

	query := r.db.Model(&model.Plan{})

	// 应用筛选条件
	for key, value := range filters {
		query = query.Where(key, value)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 应用排序
	if orderBy != "" {
		query = query.Order(orderBy)
	} else {
		query = query.Order("created_at DESC")
	}

	// 应用分页
	if err := query.Offset(offset).Limit(limit).Find(&plans).Error; err != nil {
		return nil, 0, err
	}

	return plans, total, nil
}

// Update 更新计划
func (r *planRepository) Update(plan *model.Plan) error {
	return r.db.Save(plan).Error
}

// Delete 删除计划(软删除)
func (r *planRepository) Delete(id uint) error {
	return r.db.Delete(&model.Plan{}, id).Error
}

// Count 统计计划数量
func (r *planRepository) Count(filters map[string]interface{}) (int64, error) {
	var count int64
	query := r.db.Model(&model.Plan{})

	// 应用筛选条件
	for key, value := range filters {
		query = query.Where(key, value)
	}

	if err := query.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
