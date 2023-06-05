package router

import (
	"hitss/internal/handler/user"
	"hitss/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func User(r *mux.Router) {
	handler := user.New()
	route := r.PathPrefix("/api/v1/users").Subrouter()
	route.Use(middleware.TOKEN)
	route.HandleFunc("", handler.GetAllUsers).Methods(http.MethodGet)
	route.HandleFunc("", handler.CreateUser).Methods(http.MethodPost)
	route.HandleFunc("/{id}", handler.UpdateUser).Methods(http.MethodPut)
	route.HandleFunc("/{id}", handler.DeleteUser).Methods(http.MethodDelete)
}
