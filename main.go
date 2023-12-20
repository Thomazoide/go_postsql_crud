package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Thomazoide/go_postsql_crud/routes"

	"github.com/Thomazoide/go_postsql_crud/database"
)

func main() {
	database.DBConnect()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	http.ListenAndServe(":3000", r)
}
