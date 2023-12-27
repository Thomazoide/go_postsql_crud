package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Thomazoide/go_postsql_crud/database"
	"github.com/Thomazoide/go_postsql_crud/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	database.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado..."))
	}
	json.NewEncoder(w).Encode(&user)
}
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user *models.User
	json.NewDecoder(r.Body).Decode(user)
	newUser := database.DB.Create(user)
	err := newUser.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&user)
}
func DelUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	database.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado..."))
	}
	database.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
