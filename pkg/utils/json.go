package utils

import (
	"encoding/json"
	"net/http"

	"github.com/thecodingmontana/jwt-go/pkg/types"
)

func ResponseWithJSON(res http.ResponseWriter, status int, payload interface{}) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)
	json.NewEncoder(res).Encode(payload)
}

func ResponseWithError(res http.ResponseWriter, status int, message string) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	response := types.APIResponse{
		StatusCode:    status,
		StatusMessage: message,
	}

	json.NewEncoder(res).Encode(response)
}
