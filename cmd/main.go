package main

import (
	"hitss/internal/router"
	"hitss/pkg/helper/cert"
	"hitss/pkg/helper/env"
	"hitss/pkg/middleware"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	env.Load()
	cert.Load()

	r := mux.NewRouter()
	r.Use(middleware.CORS)
	r.Use(middleware.RECOVERY)

	router.Swagger(r)
	router.Login(r)
	router.User(r)

	log.Printf("running app :%s", os.Getenv("PORT"))
	http.ListenAndServe(":"+os.Getenv("PORT"), r)
}
