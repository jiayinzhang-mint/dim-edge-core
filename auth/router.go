package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitAuthAPI init auth REST api
func InitAuthAPI(r *mux.Router) {
	privateRouter := r.PathPrefix("/api/auth").Subrouter()
	privateRouter.HandleFunc("/login", handleLogin).Methods("GET")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	return
}
