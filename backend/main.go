package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func init() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, using system environment variables")
	}

	// Setup logger
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize router
	router := mux.NewRouter()

	// Setup CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Health check endpoint
	router.HandleFunc("/health", healthCheck).Methods("GET")

	// API v1 routes
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	// Auth endpoints
	apiV1.HandleFunc("/auth/register", registerHandler).Methods("POST")
	apiV1.HandleFunc("/auth/login", loginHandler).Methods("POST")

	// Order endpoints
	apiV1.HandleFunc("/orders", listOrders).Methods("GET")
	apiV1.HandleFunc("/orders", createOrder).Methods("POST")
	apiV1.HandleFunc("/orders/{id}", getOrder).Methods("GET")

	// Payment endpoints
	apiV1.HandleFunc("/payments", createPayment).Methods("POST")
	apiV1.HandleFunc("/payments/{id}", getPayment).Methods("GET")

	// Webhook endpoints
	apiV1.HandleFunc("/webhooks/payment", paymentWebhook).Methods("POST")

	// Plugin endpoints
	apiV1.HandleFunc("/plugins", listPlugins).Methods("GET")

	// Start server
	logger.Infof("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, corsHandler(router)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// Health check
func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"status":"ok","service":"freedomth-backend"}`)
}

// Auth endpoints
func registerHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Register endpoint - Not implemented yet"}`)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Login endpoint - Not implemented yet"}`)
}

// Order endpoints
func listOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"List orders - Not implemented yet"}`)
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Create order - Not implemented yet"}`)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Get order - Not implemented yet"}`)
}

// Payment endpoints
func createPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Create payment - Not implemented yet"}`)
}

func getPayment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Get payment - Not implemented yet"}`)
}

// Webhook endpoint
func paymentWebhook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, `{"message":"Payment webhook - Not implemented yet"}`)
}

// Plugin endpoints
func listPlugins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{
		"plugins": [
			{"name":"web","status":"active"},
			{"name":"facebook","status":"active"},
			{"name":"line","status":"active"},
			{"name":"telegram","status":"active"},
			{"name":"wechat","status":"active"}
		]
	}`)
}
