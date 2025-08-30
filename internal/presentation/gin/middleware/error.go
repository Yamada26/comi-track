package middleware

import (
	"log"
	"net/http"

	"comi-track/internal/domain"

	"github.com/gin-gonic/gin"
)
func ErrorHandler() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        ctx.Next()

        if len(ctx.Errors) > 0 {
            err := ctx.Errors.Last().Err

            // 共通エラー型の場合
            if appErr, ok := err.(*domain.AppError); ok {
                status := mapKindToHTTPStatus(appErr.Kind)

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
                "error":   domain.ErrInternal.String(),
                "message": "unexpected error occurred",
            })
        }
    }
}

func mapKindToHTTPStatus(kind domain.ErrorKind) int {
    switch kind {
	case domain.ErrInvalid: // 400
		return http.StatusBadRequest
	case domain.ErrPermission: // 403
		return http.StatusForbidden
	case domain.ErrNotFound: // 404
		return http.StatusNotFound
	case domain.ErrConflict: // 409
		return http.StatusConflict
	case domain.ErrInternal: // 500
		return http.StatusInternalServerError
    default:
        return http.StatusInternalServerError
    }
}