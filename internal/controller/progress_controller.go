package controller

import (
	"strconv"
	"worker_plan/internal/service"

	"github.com/gin-gonic/gin"
)

// ProgressController 进度控制器
type ProgressController struct {
	progressService *service.ProgressService
}

// NewProgressController 创建进度控制器
func NewProgressController(progressService *service.ProgressService) *ProgressController {
	return &ProgressController{
		progressService: progressService,
	}
}

// UpdateStatus 更新状态
// @Summary 更新计划状态
// @Description 更新工作计划的状态
// @Tags progress
// @Accept json
// @Produce json
// @Param id path int true "计划 ID"
// @Param request body service.UpdateStatusRequest true "状态信息"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/plans/{id}/status [patch]
func (pc *ProgressController) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	var req service.UpdateStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequestResponse(c, err.Error())
		return
	}

	plan, err := pc.progressService.UpdateStatus(uint(id), &req)
	if err != nil {
		if err.Error() == "plan not found" {
			NotFoundResponse(c, "plan not found")
			return
		}
		BadRequestResponse(c, err.Error())
		return
	}

	SuccessResponse(c, plan)
}

// UpdateProgress 更新进度
// @Summary 更新计划进度
// @Description 更新工作计划的进度
// @Tags progress
// @Accept json
// @Produce json
// @Param id path int true "计划 ID"
// @Param request body service.UpdateProgressRequest true "进度信息"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/plans/{id}/progress [patch]
func (pc *ProgressController) UpdateProgress(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	var req service.UpdateProgressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequestResponse(c, err.Error())
		return
	}

	plan, err := pc.progressService.UpdateProgress(uint(id), &req)
	if err != nil {
		if err.Error() == "plan not found" {
			NotFoundResponse(c, "plan not found")
			return
		}
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, plan)
}
