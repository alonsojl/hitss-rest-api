package middleware

import (
	"context"
	"hitss/pkg/helper/result"
	"hitss/pkg/helper/token"
	"net/http"
	"strings"
)

func TOKEN(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := r.Header.Get("Authorization")
		if !strings.Contains(bearerToken, "Bearer ") {
			result.JSON(w, result.D{
				"code":  http.StatusBadRequest,
				"error": "bearer token not include",
			})
			return
		}
		jwtToken := strings.Trim(strings.Replace(bearerToken, "Bearer ", "", 1), " ")
		user, err := token.Validate(jwtToken)
		if err != nil {
			msg, httpCode := token.CheckError(err)
			result.JSON(w, result.D{
				"code":  httpCode,
				"error": msg,
			})
			return
		}

		ctx := r.Context()
		ctxUser := context.WithValue(ctx, result.CtxKey{}, user)
		next.ServeHTTP(w, r.WithContext(ctxUser))
	})
}
