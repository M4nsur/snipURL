package auth

import (
	"log"
	"net/http"

	config "github.com/M4nsur/snipURL/configs"
	"github.com/M4nsur/snipURL/pkg/response"
)
type AuthDeps struct { 
	*config.Config
}
type AuthHandler struct { *config.Config}

func NewAuthHandler(router *http.ServeMux, conf *config.Config) {
	handler := &AuthHandler{Config: conf}
	router.HandleFunc("POST /auth/login", handler.Register())
	router.HandleFunc("POST /auth/register", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("register its work")
		res := LoginResponse{
			Token: "123",
		}
		response.JSON(w, res, 200)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login its work")
	}
}