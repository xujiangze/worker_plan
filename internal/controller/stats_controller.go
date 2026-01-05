package controller

import (
	"worker_plan/internal/service"

	"github.com/gin-gonic/gin"
)

// StatisticsController 统计控制器
type StatisticsController struct {
	statsService *service.StatisticsService
}

// NewStatisticsController 创建统计控制器
func NewStatisticsController(statsService *service.StatisticsService) *StatisticsController {
	return &StatisticsController{
		statsService: statsService,
	}
}

// GetStatsByStatus 按状态统计
// @Summary 按状态统计
// @Description 统计各状态计划的数量和占比
// @Tags statistics
// @Produce json
// @Success 200 {object} Response
// @Router /api/stats/by-status [get]
func (sc *StatisticsController) GetStatsByStatus(c *gin.Context) {
	stats, err := sc.statsService.GetStatsByStatus()
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, stats)
}

// GetStatsByPriority 按优先级统计
// @Summary 按优先级统计
// @Description 统计各优先级计划的数量和占比
// @Tags statistics
// @Produce json
// @Success 200 {object} Response
// @Router /api/stats/by-priority [get]
func (sc *StatisticsController) GetStatsByPriority(c *gin.Context) {
	stats, err := sc.statsService.GetStatsByPriority()
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, stats)
}

// GetStatsByTime 按时间统计
// @Summary 按时间统计
// @Description 统计指定时间范围内的计划数量
// @Tags statistics
// @Produce json
// @Param start_date query string false "开始日期"
// @Param end_date query string false "结束日期"
// @Success 200 {object} Response
// @Router /api/stats/by-time [get]
func (sc *StatisticsController) GetStatsByTime(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	stats, err := sc.statsService.GetStatsByTime(startDate, endDate)
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, stats)
}

// GetCompletionRate 获取完成率
// @Summary 获取完成率
// @Description 获取计划完成率统计
// @Tags statistics
// @Produce json
// @Success 200 {object} Response
// @Router /api/stats/completion-rate [get]
func (sc *StatisticsController) GetCompletionRate(c *gin.Context) {
	rate, err := sc.statsService.GetCompletionRate()
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, rate)
}
