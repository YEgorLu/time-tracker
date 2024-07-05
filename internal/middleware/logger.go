package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/YEgorLu/time-tracker/internal/logger"
)

func Logger(log logger.Logger) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(time.Now(), r.Method, r.URL)
			next.ServeHTTP(w, r)
		}
	}
}
