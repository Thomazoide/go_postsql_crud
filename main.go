package main

import (
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/db"
	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/Thomazoide/go_postsql_crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.ConectarDB()
	DB.AutoMigrate(&models.Usuario{}, &models.Product{})
	router := mux.NewRouter()
	Handler := routes.NewHandler(DB)
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users/{id}", Handler.GetUser).Methods("GET")
	router.HandleFunc("/users", Handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", Handler.SignUp).Methods("POST")
	http.ListenAndServe(":3132", router)
}
