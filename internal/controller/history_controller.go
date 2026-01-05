package controller

import (
	"strconv"
	"worker_plan/internal/service"

	"github.com/gin-gonic/gin"
)

// HistoryController 历史控制器
type HistoryController struct {
	historyService *service.HistoryService
}

// NewHistoryController 创建历史控制器
func NewHistoryController(historyService *service.HistoryService) *HistoryController {
	return &HistoryController{
		historyService: historyService,
	}
}

// GetHistory 获取历史记录
// @Summary 获取计划历史记录
// @Description 获取指定计划的历史变更记录
// @Tags history
// @Produce json
// @Param id path int true "计划 ID"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/plans/{id}/history [get]
func (hc *HistoryController) GetHistory(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))

	histories, total, err := hc.historyService.GetHistoryByPlanID(uint(id), page, pageSize)
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, gin.H{
		"total":     total,
		"page":      page,
		"page_size": pageSize,
		"data":      histories,
	})
}
