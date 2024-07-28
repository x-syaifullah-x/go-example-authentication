package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/x-syaifullah-x/go-crud/src/pkg/logger"
)

type authMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *authMiddleware {
	return &authMiddleware{handler: handler}
}

func (middleware *authMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf(
		"\n\tRequest From : %s\n\tEnd Point    : %s%s\n\tMethod       : %s",
		r.RemoteAddr,
		r.Host,
		r.RequestURI,
		r.Method,
	)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if "RAHASIA" == r.Header.Get("X-API-Key") {
		middleware.handler.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		response := map[string]any{
			"code":   http.StatusUnauthorized,
			"status": "UNAUTHORIZED",
		}
		results, _ := json.Marshal(&response)
		w.Write(results)
	}
}
