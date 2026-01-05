package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"worker_plan/internal/config"
	"worker_plan/internal/controller"
	"worker_plan/internal/middleware"
	"worker_plan/internal/model"
	"worker_plan/internal/repository"
	"worker_plan/internal/router"
	"worker_plan/internal/service"
	"worker_plan/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 初始化日志
	if err := middleware.InitLogger(cfg.Log.Level, cfg.Log.Format); err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	middleware.InfoLog("Starting worker plan server...")

	// 初始化数据库
	dbConfig := database.DefaultConfig(cfg.Database.GetDSN())
	if err := database.Init(dbConfig); err != nil {
		middleware.ErrorLog("Failed to init database")
		log.Fatalf("Failed to init database: %v", err)
	}
	defer database.Close()

	middleware.InfoLog("Database connected successfully")

	// 自动迁移数据库表
	if err := autoMigrate(); err != nil {
		middleware.ErrorLog("Failed to migrate database")
		log.Fatalf("Failed to migrate database: %v", err)
	}

	middleware.InfoLog("Database migration completed")

	// 初始化仓储
	repos := repository.NewRepositories(database.GetDB())

	// 初始化服务
	planService := service.NewPlanService(repos.Plan, repos.PlanHistory)
	progressService := service.NewProgressService(repos.Plan, repos.PlanHistory)
	statsService := service.NewStatisticsService(repos.Plan)
	historyService := service.NewHistoryService(repos.PlanHistory)

	// 初始化控制器
	planController := controller.NewPlanController(planService)
	progressController := controller.NewProgressController(progressService)
	statsController := controller.NewStatisticsController(statsService)
	historyController := controller.NewHistoryController(historyService)

	// 设置路由
	router := router.SetupRouter(
		planController,
		progressController,
		statsController,
		historyController,
	)

	// 设置 Gin 模式
	gin.SetMode(cfg.Server.Mode)

	// 创建 HTTP 服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// 启动服务器
	go func() {
		middleware.InfoLog(fmt.Sprintf("Server is running on %s", addr))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			middleware.ErrorLog("Failed to start server")
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	middleware.InfoLog("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		middleware.ErrorLog("Server forced to shutdown")
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	middleware.InfoLog("Server exited")
}

// autoMigrate 自动迁移数据库表
func autoMigrate() error {
	db := database.GetDB()

	// 迁移 Plan 模型
	if err := db.AutoMigrate(&model.Plan{}); err != nil {
		return fmt.Errorf("failed to migrate Plan model: %w", err)
	}

	// 迁移 PlanHistory 模型
	if err := db.AutoMigrate(&model.PlanHistory{}); err != nil {
		return fmt.Errorf("failed to migrate PlanHistory model: %w", err)
	}

	return nil
}
