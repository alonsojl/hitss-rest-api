package middleware

import (
	"hitss/pkg/helper/logger"
	"hitss/pkg/helper/result"
	"net/http"
)

func RECOVERY(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Write(err)
				result.JSON(w, result.D{
					"code":  http.StatusInternalServerError,
					"error": "internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
