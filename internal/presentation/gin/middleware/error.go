package middleware

import (
	"log"
	"net/http"

	"comi-track/internal/common"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			// 共通エラー型の場合
			if appErr, ok := err.(*common.AppError); ok {
				status := appErr.Kind.StatusCode()

				// レスポンスには安全な情報だけ返す
				ctx.JSON(status, gin.H{
					"error":   appErr.Kind.String(),
					"message": appErr.Message,
				})
				return
			}

			// 自作でないエラー
			log.Printf("[PANIC] %v", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error":   common.ErrInternal.String(),
				"message": "unexpected error occurred",
			})
		}
	}
}
