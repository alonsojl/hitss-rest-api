package router

import (
	"hitss/internal/handler/swagger"
	"net/http"

	"github.com/gorilla/mux"
)

func Swagger(r *mux.Router) {
	handler := swagger.New()
	r.PathPrefix("/ui").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./asset/swagger/"))))
	r.HandleFunc("/openapi3.json", handler.GetJsonOpenAPI).Methods(http.MethodGet)
	r.HandleFunc("/openapi3.yaml", handler.GetYamlOpenAPI).Methods(http.MethodGet)
}
