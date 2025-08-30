package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        start := time.Now()
        ctx.Next() // ハンドラ実行
        latency := time.Since(start)

        logger.Info("Request completed",
            "method", ctx.Request.Method,
            "path", ctx.Request.URL.Path,
            "status", ctx.Writer.Status(),
            "latency", latency,
            "client_ip", ctx.ClientIP(),
        )
    }
}