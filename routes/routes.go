package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRoutes() http.Handler {
	router := chi.NewRouter()
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Gateway is running"))
	})

	router.Get("/manifest", ManifestHandler)

	router.Post("/api/webhook/order", WebhookHandler)

	return router
}
