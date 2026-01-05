package router

import (
	"worker_plan/internal/controller"
	"worker_plan/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(
	planController *controller.PlanController,
	progressController *controller.ProgressController,
	statsController *controller.StatisticsController,
	historyController *controller.HistoryController,
) *gin.Engine {
	router := gin.Default()

	// 应用中间件
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RecoveryMiddleware())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// API 路由组
	api := router.Group("/api")
	{
		// 计划管理路由
		plans := api.Group("/plans")
		{
			plans.POST("", planController.CreatePlan)
			plans.GET("", planController.GetPlans)
			plans.GET("/:id", planController.GetPlan)
			plans.PUT("/:id", planController.UpdatePlan)
			plans.DELETE("/:id", planController.DeletePlan)

			// 进度管理路由
			plans.PATCH("/:id/status", progressController.UpdateStatus)
			plans.PATCH("/:id/progress", progressController.UpdateProgress)

			// 历史记录路由
			plans.GET("/:id/history", historyController.GetHistory)
		}

		// 统计分析路由
		stats := api.Group("/stats")
		{
			stats.GET("/by-status", statsController.GetStatsByStatus)
			stats.GET("/by-priority", statsController.GetStatsByPriority)
			stats.GET("/by-time", statsController.GetStatsByTime)
			stats.GET("/completion-rate", statsController.GetCompletionRate)
		}
	}

	return router
}
