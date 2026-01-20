package auth

import (
	"fmt"
	"net/http"
)

type AuthHandler struct{}


func NewAuthHandlers() * AuthHandler{
	return &AuthHandler{}
}


func (h *AuthHandler) Login (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func (h *AuthHandler) Register (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}