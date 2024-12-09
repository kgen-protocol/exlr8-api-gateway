package routes

import (
	"encoding/json"
	"net/http"
)

type Manifest struct {
	ID               string    `json:"id"`
	Version          string    `json:"version"`
	Name             string    `json:"name"`
	Permissions      []string  `json:"permissions"`
	AppURL           string    `json:"appUrl"`
	ConfigurationURL string    `json:"configurationUrl"`
	TokenTargetURL   string    `json:"tokenTargetUrl"`
	Webhooks         []Webhook `json:"webhooks"`
}

type Webhook struct {
	Name      string `json:"name"`
	EventType string `json:"eventType"`
	TargetURL string `json:"targetUrl"`
}

func ManifestHandler(w http.ResponseWriter, r *http.Request) {
	manifest := Manifest{
		ID:      "api-gateway",
		Version: "1.0",
		Name:    "Go API Gateway",
		Permissions: []string{
			"MANAGE_ORDERS",
			"MANAGE_PRODUCTS",
		},
		AppURL:           "http://localhost:8080/",
		ConfigurationURL: "http://localhost:8080/configure",
		TokenTargetURL:   "http://localhost:8080/register-token",
		Webhooks:         []Webhook{
			// {
			// 	Name:      "Order Created",
			// 	EventType: "ORDER_CREATED",
			// 	TargetURL: "https://ffc6-2401-4900-1cbd-92f5-95fb-df19-f55a-62cd.ngrok-free.app/api/webhook/order",
			// },
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(manifest)
}
