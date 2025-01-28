package middlewares

import "net/http"

const MaxRequestBodySize = 1 * 1024 * 1024 // 1 MB

func LimitRequestBodySize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, MaxRequestBodySize)

		next.ServeHTTP(w, r)
	})
}
