package handlers

import (
	"net/http"

	"github.com/thecodingmontana/jwt-go/pkg/types"
	"github.com/thecodingmontana/jwt-go/pkg/utils"
)

func (api *APIConfig) CreateUserHandler(res http.ResponseWriter, req *http.Request) {
	utils.ResponseWithJSON(res, http.StatusOK, types.APIResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "Hello user create route",
	})
}
