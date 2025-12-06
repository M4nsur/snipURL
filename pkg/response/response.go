package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, data any, resCode int) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resCode)
		json.NewEncoder(w).Encode(data) 
}