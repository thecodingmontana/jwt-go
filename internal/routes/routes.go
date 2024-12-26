package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/thecodingmontana/jwt-go/pkg/types"
	"github.com/thecodingmontana/jwt-go/pkg/utils"
)

func RegisterRoutes(router chi.Router) {
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		utils.ResponseWithJSON(res, http.StatusOK, types.APIResponse{
			StatusCode:    http.StatusOK,
			StatusMessage: "Welcome to JWT-GO API!",
		})
	})

	router.Get("/healthz", func(res http.ResponseWriter, req *http.Request) {
		utils.ResponseWithJSON(res, http.StatusOK, types.APIResponse{
			StatusCode:    http.StatusOK,
			StatusMessage: "Everthing working correctly!",
		})
	})
}
