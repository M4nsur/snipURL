package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/M4nsur/snipURL/configs"
)

type RegisterResponse struct {
    Message string `json:"message"`
    UserID  int    `json:"user_id"`
}


type AuthHandler struct{
	*configs.Config
}

func NewAuthHandlers(deps *configs.Config) *AuthHandler{
	return &AuthHandler{
		deps,
	}
}

func (h *AuthHandler) Login (w http.ResponseWriter, r *http.Request) {		
	var loginData LoginDTO
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
	respondWithJSON(w, http.StatusOK, "mockToken")
}

func (h *AuthHandler) Register (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")
	respondWithJSON(w, http.StatusOK, RegisterResponse{
		Message: "user registed",
		UserID: 21389,
	})

}



func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, `{"message":"failed to encode response"}`, http.StatusInternalServerError)
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	errDTO := ErrorDTO{
		Message: message,
		Time:    time.Now(),
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errDTO)
}

