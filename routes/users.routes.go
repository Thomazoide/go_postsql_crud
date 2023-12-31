package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.Usuario
	h.DB.First(&user, id)
	if user.ID == 0 {
		var err SvResponse
		err.Mensaje = "Usuario no encontrado..."
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(user)
	}
}

func (h *handler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	h.DB.Find(&usuarios)
	if len(usuarios) == 0 {
		var res SvResponse
		res.Mensaje = "No hay usuarios registrados..."
		json.NewEncoder(w).Encode(res)
	} else {
		json.NewEncoder(w).Encode(usuarios)
	}
}

func (h *handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.Usuario
	json.NewDecoder(r.Body).Decode(&user)
	h.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.Usuario
	h.DB.First(&user, id)
	if user.ID == 0 {
		var err SvResponse
		err.Mensaje = "Usuario ya no se encuentra en la base de datos..."
		json.NewEncoder(w).Encode(err)
	} else {
		var res SvResponse
		res.Mensaje = "Usuario borrado..."
		h.DB.Delete(&user)
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(res)
	}
}

func (h *handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var user models.Usuario
	h.DB.First(&user, id)
	if user.ID == 0 {
		var err SvResponse
		err.Mensaje = "Error en la operacion..."
		log.Fatal("No encontrado")
		json.NewEncoder(w).Encode(err)
	}
}
