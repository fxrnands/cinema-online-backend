package routes

import (
	"cinema-online/handlers"
	"cinema-online/pkg/middleware"
	"cinema-online/pkg/mysql"
	"cinema-online/repositories"

	"github.com/gorilla/mux"
)

func CategoryRoutes(r *mux.Router) {
	categoryRepository := repositories.RepositoryCategory(mysql.DB)
	h := handlers.HandlerCategory(categoryRepository)

	r.HandleFunc("/categorys", h.FindCategorys).Methods("GET")
	r.HandleFunc("/category/{id}", h.GetCategory).Methods("GET")
	r.HandleFunc("/category", middleware.Auth(h.CreateCategory)).Methods("POST")
	r.HandleFunc("/category/{id}", middleware.Auth(h.UpdateCategory)).Methods("PATCH")
	r.HandleFunc("/category/{id}", middleware.Auth(h.DeleteCategory)).Methods("DELETE")
}
