package response

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, data any, resCode int) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(resCode)
		err := json.NewEncoder(w).Encode(data) 
		if err != nil {
			log.Printf("Failed to encode JSON: %v", err)
		}
}