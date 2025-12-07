package auth

import (
	"net/http"

	config "github.com/M4nsur/snipURL/configs"
	"github.com/M4nsur/snipURL/pkg/request"
	"github.com/M4nsur/snipURL/pkg/response"
)
type AuthDeps struct { 
	*config.Config
}
type AuthHandler struct { *config.Config}

func NewAuthHandler(router *http.ServeMux, conf *config.Config) {
	handler := &AuthHandler{Config: conf}
	router.HandleFunc("POST /auth/register", handler.Register())
	router.HandleFunc("POST /auth/login", handler.Login())
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}
		res := AuthResponse{Token: "123",
		}
		response.JSON(w, res, 200)
	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}
		res := AuthResponse{Token: "123",
		}
		response.JSON(w, res, 200)
	}
}