package service

import (
	"fmt"
	"time"
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
	CreatedCount   int64             `json:"created_count"`
	CompletedCount int64             `json:"completed_count"`
	CompletionRate float64           `json:"completion_rate"`
	DailyTrend     []*DailyTrendItem `json:"daily_trend"`
}

// DailyTrendItem 每日趋势项
type DailyTrendItem struct {
	Date      string `json:"date"`
	Created   int64  `json:"created"`
	Completed int64  `json:"completed"`
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

// GetStatsByTime 按时间统计
func (s *StatisticsService) GetStatsByTime(startDate, endDate string) (*TimeStats, error) {
	// 验证日期格式
	if startDate != "" {
		if _, err := time.Parse("2006-01-02", startDate); err != nil {
			return nil, fmt.Errorf("invalid start date format: %w", err)
		}
	}
	if endDate != "" {
		if _, err := time.Parse("2006-01-02", endDate); err != nil {
			return nil, fmt.Errorf("invalid end date format: %w", err)
		}
	}

	// 验证日期范围
	if startDate != "" && endDate != "" {
		start, _ := time.Parse("2006-01-02", startDate)
		end, _ := time.Parse("2006-01-02", endDate)
		if start.After(end) {
			return nil, fmt.Errorf("start date cannot be after end date")
		}
	}

	// 获取创建数量
	createdCount, err := s.planRepo.CountByDateRange(startDate, endDate, "")
	if err != nil {
		return nil, fmt.Errorf("failed to get created count: %w", err)
	}

	// 获取完成数量
	completedCount, err := s.planRepo.CountByDateRange(startDate, endDate, "Done")
	if err != nil {
		return nil, fmt.Errorf("failed to get completed count: %w", err)
	}

	// 计算完成率
	completionRate := 0.0
	if createdCount > 0 {
		completionRate = float64(completedCount) / float64(createdCount) * 100
	}

	// 获取每日趋势数据
	dailyTrend, err := s.planRepo.GetDailyTrend(startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("failed to get daily trend: %w", err)
	}

	// 转换为service层的DailyTrendItem
	var trendItems []*DailyTrendItem
	for _, item := range dailyTrend {
		trendItems = append(trendItems, &DailyTrendItem{
			Date:      item.Date,
			Created:   item.Created,
			Completed: item.Completed,
		})
	}

	return &TimeStats{
		CreatedCount:   createdCount,
		CompletedCount: completedCount,
		CompletionRate: completionRate,
		DailyTrend:     trendItems,
	}, nil
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
