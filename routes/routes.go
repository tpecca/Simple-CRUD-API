package routes

import (
	"CRUD-API/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("A very simple API with Golang"))
	})

	r.Get("/item", handlers.GetHandler)
	r.Post("/item", handlers.PostHandler)
	r.Put("/item", handlers.PutHandler)
	r.Delete("/item", handlers.DeleteHandler)

	return r
}
