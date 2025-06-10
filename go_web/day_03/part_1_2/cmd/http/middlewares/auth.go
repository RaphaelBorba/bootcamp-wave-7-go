package middlewares

import (
	"net/http"

	"github.com/bootcamp-go/web/response"
)

// Auth verifica se o header Authorization está presente e é válido.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || token != "123456abc" {
			response.JSON(w, http.StatusUnauthorized, "Authorization token obrigatório")
			return
		}

		next.ServeHTTP(w, r)
	})
}
