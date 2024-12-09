package routes

import (
	"fmt"
	"io"
	"net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Log or process webhook payload
	fmt.Printf("Received Webhook: %s\n", string(body))
	w.WriteHeader(http.StatusOK)
}
