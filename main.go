package main

import (
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/db"
	"github.com/Thomazoide/go_postsql_crud/middlewares"
	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/Thomazoide/go_postsql_crud/routes"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.ConectarDB()
	DB.AutoMigrate(&models.Usuario{}, &models.Product{}, &models.Cart{})
	router := mux.NewRouter()
	Handler := routes.NewHandler(DB)
	middlewares.EnableCORS(router)
	router.Use(middlewares.MiddlewareCORS)
	router.HandleFunc("/", routes.HomeHandler)
	router.HandleFunc("/users/{id}", Handler.GetUser).Methods("GET")
	router.HandleFunc("/users", Handler.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", Handler.SignUp).Methods("POST")
	router.HandleFunc("/login", Handler.LogIn).Methods("POST")
	router.HandleFunc("/login", Handler.VerifyToken).Methods("PUT")
	router.HandleFunc("/users/{id}/cart/{pid}", Handler.AddProduct).Methods("PUT")
	router.HandleFunc("/products", Handler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", Handler.GetProduct).Methods("GET")
	router.HandleFunc("/products", Handler.NewProduct).Methods("POST")
	http.ListenAndServe(":3132", router)
}
