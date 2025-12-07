package request

import (
	"net/http"

	"github.com/M4nsur/snipURL/pkg/response"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request)(*T, error)  {
	body, err := Decode[T](r.Body)
	if err != nil {
		response.JSON(w, err.Error(), 400)
		return nil, err
	}
	err = isValid(body) 
	if err != nil {
		response.JSON(w, err.Error(), 400)
		return nil, err
	}
	return &body, nil
}