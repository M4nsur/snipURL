package auth

import (
	"fmt"
	"net/http"

	"github.com/M4nsur/snipURL/configs"
)

type AuthHandler struct{
	*configs.Config
}

func NewAuthHandlers(deps *configs.Config) *AuthHandler{
	return &AuthHandler{
		deps,
	}
}

func (h *AuthHandler) Login (w http.ResponseWriter, r *http.Request) {	
	fmt.Println("Login")
	fmt.Println(h.Config.Auth.Secret, h.Config.Db.DNS)
}

func (h *AuthHandler) Register (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
}