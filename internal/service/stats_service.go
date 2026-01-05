package service

import (
	"fmt"
	"worker_plan/internal/repository"
)

// StatisticsService 统计服务
type StatisticsService struct {
	planRepo repository.PlanRepository
}

// NewStatisticsService 创建统计服务
func NewStatisticsService(planRepo repository.PlanRepository) *StatisticsService {
	return &StatisticsService{
		planRepo: planRepo,
	}
}

// StatusStats 状态统计
type StatusStats struct {
	Status  string  `json:"status"`
	Count   int64   `json:"count"`
	Percent float64 `json:"percent"`
}

// PriorityStats 优先级统计
type PriorityStats struct {
	Priority string  `json:"priority"`
	Count    int64   `json:"count"`
	Percent  float64 `json:"percent"`
}

// TimeStats 时间统计
type TimeStats struct {
	Date  string `json:"date"`
	Count int64  `json:"count"`
}

// CompletionRate 完成率
type CompletionRate struct {
	TotalPlans     int64   `json:"total_plans"`
	CompletedPlans int64   `json:"completed_plans"`
	CompletionRate float64 `json:"completion_rate"`
}

// GetStatsByStatus 按状态统计
func (s *StatisticsService) GetStatsByStatus() ([]*StatusStats, error) {
	statuses := []string{"Todo", "InProgress", "Done", "Cancelled"}
	var stats []*StatusStats

	// 获取总数
	total, err := s.planRepo.Count(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	// 统计各状态数量
	for _, status := range statuses {
		count, err := s.planRepo.Count(map[string]interface{}{"status = ?": status})
		if err != nil {
			return nil, fmt.Errorf("failed to get count for status %s: %w", status, err)
		}

		percent := 0.0
		if total > 0 {
			percent = float64(count) / float64(total) * 100
		}

		stats = append(stats, &StatusStats{
			Status:  status,
			Count:   count,
			Percent: percent,
		})
	}

	return stats, nil
}

// GetStatsByPriority 按优先级统计
func (s *StatisticsService) GetStatsByPriority() ([]*PriorityStats, error) {
	priorities := []string{"High", "Medium", "Low"}
	var stats []*PriorityStats

	// 获取总数
	total, err := s.planRepo.Count(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	// 统计各优先级数量
	for _, priority := range priorities {
		count, err := s.planRepo.Count(map[string]interface{}{"priority = ?": priority})
		if err != nil {
			return nil, fmt.Errorf("failed to get count for priority %s: %w", priority, err)
		}

		percent := 0.0
		if total > 0 {
			percent = float64(count) / float64(total) * 100
		}

		stats = append(stats, &PriorityStats{
			Priority: priority,
			Count:    count,
			Percent:  percent,
		})
	}

	return stats, nil
}

// GetStatsByTime 按时间统计(简化版,按天统计)
func (s *StatisticsService) GetStatsByTime(startDate, endDate string) ([]*TimeStats, error) {
	// 这里简化处理,实际应该解析日期并使用 SQL 的 DATE_TRUNC 函数
	// 由于 repository 层的限制,这里返回空数组
	// 实际实现需要在 repository 层添加更复杂的查询方法

	var stats []*TimeStats

	// TODO: 实现按时间统计的逻辑
	// 需要在 repository 层添加支持日期范围统计的方法

	return stats, nil
}

// GetCompletionRate 获取完成率
func (s *StatisticsService) GetCompletionRate() (*CompletionRate, error) {
	// 获取总数
	total, err := s.planRepo.Count(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	// 获取已完成数量
	completed, err := s.planRepo.Count(map[string]interface{}{"status = ?": "Done"})
	if err != nil {
		return nil, fmt.Errorf("failed to get completed count: %w", err)
	}

	completionRate := 0.0
	if total > 0 {
		completionRate = float64(completed) / float64(total) * 100
	}

	return &CompletionRate{
		TotalPlans:     total,
		CompletedPlans: completed,
		CompletionRate: completionRate,
	}, nil
}
