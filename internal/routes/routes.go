package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router) {
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Hello World!"))
	})
}
