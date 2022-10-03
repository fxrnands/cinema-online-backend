package routes

import (
	"cinema-online/handlers"
	"cinema-online/pkg/middleware"
	"cinema-online/pkg/mysql"
	"cinema-online/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
  authRepository := repositories.RepositoryAuth(mysql.DB)
  h := handlers.HandlerAuth(authRepository)

  r.HandleFunc("/register", h.Register).Methods("POST")
  r.HandleFunc("/login", h.Login).Methods("POST")
  r.HandleFunc("/check-auth", middleware.Auth(h.CheckAuth)).Methods("GET")
}