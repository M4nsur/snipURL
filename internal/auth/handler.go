package auth

import (
	"encoding/json"
	"log"
	"net/http"

	config "github.com/M4nsur/snipURL/configs"
	"github.com/M4nsur/snipURL/pkg/response"
	"github.com/go-playground/validator"
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
		log.Println("register its work")

	}
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload LoginRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			response.JSON(w, err.Error(), 400)
			return 
		}
		validate := validator.New()
		errValidate := validate.Struct(payload)
		if errValidate != nil {
			response.JSON(w, errValidate.Error(), 400)
			return
		}  
		res := LoginResponse{Token: "123",
		}
		response.JSON(w, res, 200)
	}
}