package controller

import (
	"A1_BMS_PROJECT/model"
	"A1_BMS_PROJECT/service"
	"encoding/json"
	"net/http"
)

type AuthController struct {
	AuthService *service.AuthService
}

func NewAuthController(service *service.AuthService) *AuthController {
	return &AuthController{AuthService: service}
}

func (c *AuthController) Signup(w http.ResponseWriter, r *http.Request) {
	var credentials model.User

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := c.AuthService.Signup(credentials.Username, credentials.Password); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}
