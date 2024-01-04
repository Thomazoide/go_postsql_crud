package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Thomazoide/go_postsql_crud/models"
)

func (h *handler) LogIn(w http.ResponseWriter, r *http.Request) {
	var users []models.Usuario
	h.DB.Find(&users)
	var req SessionPayload
	json.NewDecoder(r.Body).Decode(&req)
	var index int = -1
	for i, user := range users {
		if user.Email == req.Email {
			index = i
		}
	}
	if index == -1 {
		var res = &SvResponse{
			Mensaje: "Email no registrado...",
		}
		json.NewEncoder(w).Encode(res)
	}
	var passCheck bool = CheckPasswordHash(req.Password, users[index].Password)
	if !passCheck {
		var res = &SvResponse{
			Mensaje: "Contraseña incorrecta...",
		}
		json.NewEncoder(w).Encode(res)
	}
	usrToken, err := GenerateToken(users[index].Email)
	if err != nil {
		var res = &SvResponse{
			Mensaje: "Error al crear token...",
		}
		json.NewEncoder(w).Encode(res)
	}
	var res = &SvResponse{
		Mensaje: "Inicio de sesion exitoso...",
		Cuerpo:  usrToken,
	}
	json.NewEncoder(w).Encode(res)
}

func (h *handler) VerifyToken(w http.ResponseWriter, r *http.Request) {
	var req SessionCheck
	json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		var res = &SvResponse{
			Mensaje: "Error al leer token...",
		}
		json.NewEncoder(w).Encode(res)
	}
	claim, err := ValidateToken(req.Token)
	if err != nil {
		var res = &SvResponse{
			Mensaje: "Error al validar token",
			Cuerpo:  err,
		}
		json.NewEncoder(w).Encode(res)
	}
	var res = &SvResponse{
		Mensaje: "Token valido!",
		Cuerpo:  claim,
	}
	json.NewEncoder(w).Encode(res)
}
