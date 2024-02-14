package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/models"
	"github.com/gorilla/mux"
)

func (h *handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var producto models.Product
	h.DB.First(&producto, id)
	if producto.ID == 0 {
		var res = &SvResponse{
			Mensaje: "Producto no encontrado...",
		}
		json.NewEncoder(w).Encode(res)
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(producto)
}

func (h *handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var productos []models.Product
	h.DB.Find(&productos)
	if len(productos) == 0 {
		var res = &SvResponse{
			Mensaje: "No hay productos existentes...",
		}
		json.NewEncoder(w).Encode(res)
	} else {
		var res = &SvResponse{
			Mensaje: "Productos encontrados!",
			Cuerpo:  productos,
		}
		json.NewEncoder(w).Encode(res)
	}
}

func (h *handler) NewProduct(w http.ResponseWriter, r *http.Request) {
	var producto models.Product
	json.NewDecoder(r.Body).Decode(&producto)
	h.DB.Create(&producto)
	var res = &SvResponse{
		Mensaje: "Producto agregado!",
		Cuerpo:  producto,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var producto models.Product
	var patchProduct models.Product
	json.NewDecoder(r.Body).Decode(&patchProduct)
	h.DB.First(&producto, id)
	if producto.ID == 0 {
		var res SvResponse
		res.Mensaje = "Error en la operacion..."
		json.NewEncoder(w).Encode(res)
	} else {
		var res SvResponse
		patchProduct.ID = producto.ID
		h.DB.Save(&patchProduct)
		res.Mensaje = "Producto guardado..."
		res.Cuerpo = patchProduct
		json.NewEncoder(w).Encode(res)
	}
}
