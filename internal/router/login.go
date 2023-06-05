package router

import (
	"hitss/internal/handler/login"
	"net/http"

	"github.com/gorilla/mux"
)

func Login(r *mux.Router) {
	handler := login.New()
	r.HandleFunc("/api/v1/signin", handler.SignInUser).Methods(http.MethodPost)
}
