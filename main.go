package main

import (
	"github.com/gorilla/mux"
	"github.com/mahmudulhasandip/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

//func init() {
//  config.LoadEnv()
//}

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("go-book-crud-production.up.railway.app", r))
}
