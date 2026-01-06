package controller

import (
	"strconv"
	"worker_plan/internal/service"

	"github.com/gin-gonic/gin"
)

// PlanController 计划控制器
type PlanController struct {
	planService *service.PlanService
}

// NewPlanController 创建计划控制器
func NewPlanController(planService *service.PlanService) *PlanController {
	return &PlanController{
		planService: planService,
	}
}

// CreatePlan 创建计划
// @Summary 创建计划
// @Description 创建新的工作计划
// @Tags plans
// @Accept json
// @Produce json
// @Param request body service.CreatePlanRequest true "计划信息"
// @Success 201 {object} Response
// @Failure 400 {object} Response
// @Router /api/plans [post]
func (pc *PlanController) CreatePlan(c *gin.Context) {
	var req service.CreatePlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequestResponse(c, err.Error())
		return
	}

	plan, err := pc.planService.CreatePlan(&req)
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	CreatedResponse(c, plan)
}

// GetPlans 获取计划列表
// @Summary 获取计划列表
// @Description 获取工作计划列表,支持分页和筛选
// @Tags plans
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(20)
// @Param status query string false "状态筛选"
// @Param priority query string false "优先级筛选"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Router /api/plans [get]
func (pc *PlanController) GetPlans(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	status := c.Query("status")
	priority := c.Query("priority")

	plans, total, err := pc.planService.GetPlans(page, pageSize, status, priority)
	if err != nil {
		InternalErrorResponse(c, err.Error())
		return
	}

	SuccessResponse(c, gin.H{
		"items":       plans,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// GetPlan 获取单个计划
// @Summary 获取计划
// @Description 根据 ID 获取工作计划详情
// @Tags plans
// @Produce json
// @Param id path int true "计划 ID"
// @Success 200 {object} Response
// @Failure 404 {object} Response
// @Router /api/plans/{id} [get]
func (pc *PlanController) GetPlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	plan, err := pc.planService.GetPlan(uint(id))
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

// UpdatePlan 更新计划
// @Summary 更新计划
// @Description 更新工作计划信息
// @Tags plans
// @Accept json
// @Produce json
// @Param id path int true "计划 ID"
// @Param request body service.UpdatePlanRequest true "更新信息"
// @Success 200 {object} Response
// @Failure 400 {object} Response
// @Failure 404 {object} Response
// @Router /api/plans/{id} [put]
func (pc *PlanController) UpdatePlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	var req service.UpdatePlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		BadRequestResponse(c, err.Error())
		return
	}

	plan, err := pc.planService.UpdatePlan(uint(id), &req)
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

// DeletePlan 删除计划
// @Summary 删除计划
// @Description 删除工作计划(软删除)
// @Tags plans
// @Produce json
// @Param id path int true "计划 ID"
// @Success 204
// @Failure 404 {object} Response
// @Router /api/plans/{id} [delete]
func (pc *PlanController) DeletePlan(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		BadRequestResponse(c, "invalid plan id")
		return
	}

	if err := pc.planService.DeletePlan(uint(id)); err != nil {
		if err.Error() == "plan not found" {
			NotFoundResponse(c, "plan not found")
			return
		}
		InternalErrorResponse(c, err.Error())
		return
	}

	NoContentResponse(c)
}
