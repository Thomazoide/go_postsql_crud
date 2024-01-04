package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) AddProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	productId := vars["pid"]
	var user models.Usuario
	h.DB.First(&user, userId)
	var newProduct models.Product
	h.DB.First(&newProduct, productId)
	if user.ID == 0 {
		var err SvResponse
		err.Mensaje = "Error en la operacion..."
		log.Fatal("No encontrado")
		json.NewEncoder(w).Encode(err)
	}
	if newProduct.ID == 0 {
		err := &SvResponse{
			Mensaje: "Producto no encontrado...",
			Cuerpo:  nil,
		}
		log.Fatal("No encontrado")
		json.NewEncoder(w).Encode(err)
	}
	user.Carro.Productos = append(user.Carro.Productos, newProduct)
	h.DB.Save(&user)
	var res = &SvResponse{
		Mensaje: "Producto agregado...",
		Cuerpo:  user.Carro,
	}
	json.NewEncoder(w).Encode(res)
}
