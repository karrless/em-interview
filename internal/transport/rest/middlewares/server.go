package middlewares

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karrless/em-interview/pkg/logger"
	"go.uber.org/zap"
)

func WithLogger(ctx *context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		logger.GetLoggerFromCtx(*ctx).Info("Request",
			zap.Int("status", statusCode),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("duration", duration.String()),
		)
	}

}