package repository

import (
	"errors"
	"worker_plan/internal/model"

	"gorm.io/gorm"
)

// ErrRecordNotFound 记录未找到错误
var ErrRecordNotFound = errors.New("record not found")

// PlanRepository 计划仓储接口
type PlanRepository interface {
	Create(plan *model.Plan) error
	FindByID(id uint) (*model.Plan, error)
	FindAll(offset, limit int, filters map[string]interface{}, orderBy string) ([]*model.Plan, int64, error)
	Update(plan *model.Plan) error
	Delete(id uint) error
	Count(filters map[string]interface{}) (int64, error)
	CountByDateRange(startDate, endDate string, status string) (int64, error)
	GetDailyTrend(startDate, endDate string) ([]*DailyTrendItem, error)
}

// PlanHistoryRepository 计划历史仓储接口
type PlanHistoryRepository interface {
	Create(history *model.PlanHistory) error
	FindByPlanID(planID uint, offset, limit int) ([]*model.PlanHistory, int64, error)
	CountByPlanID(planID uint) (int64, error)
}

// Repositories 仓储集合
type Repositories struct {
	Plan        PlanRepository
	PlanHistory PlanHistoryRepository
}

// NewRepositories 创建仓储集合
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		Plan:        NewPlanRepository(db),
		PlanHistory: NewPlanHistoryRepository(db),
	}
}
