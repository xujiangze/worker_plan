package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RecoveryMiddleware 恢复中间件
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录 panic 日志
				if Logger != nil {
					Logger.Error("Panic recovered",
						zap.Any("error", err),
						zap.String("path", c.Request.URL.Path),
						zap.String("method", c.Request.Method),
					)
				}

				// 返回 500 错误
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": "Internal Server Error",
				})
				c.Abort()
			}
		}()

		c.Next()
	}
}
