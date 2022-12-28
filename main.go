package main

import (
	"github.com/gorilla/mux"
	"github.com/mahmudulhasandip/go-bookstore/pkg/routes"
	"log"
	"net/http"
	"os"
)

//func init() {
//  config.LoadEnv()
//}

func main() {
	//config.LoadEnv()
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))
}
