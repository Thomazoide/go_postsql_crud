package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Thomazoide/go_postsql_crud/routes"

	"github.com/Thomazoide/go_postsql_crud/database"

	"github.com/Thomazoide/go_postsql_crud/models"
)

func main() {
	database.DBConnect()
	database.DB.AutoMigrate(models.User{})
	database.DB.AutoMigrate(models.Product{})
	database.DB.AutoMigrate(models.Stock{})
	database.DB.AutoMigrate(models.Cart{})
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/users", routes.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUser).Methods("GET")
	r.HandleFunc("/users", routes.PostUser).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DelUser).Methods("DELETE")
	http.ListenAndServe(":3000", r)
}
