package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Extract the "action" from the request body (JSON)
func Router(w http.ResponseWriter, r *http.Request) {
	// Handle preflight request for CORS
	w.Header().Set("Allow", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return // Preflight request handled
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	bodyBytes, _ := io.ReadAll(r.Body) // Read the body to handle errors
	_ = r.Body.Close()

	var api API
	if err := json.Unmarshal(bodyBytes, &api); err != nil {
		http.Error(w, "Invalid JSON in request body or missing action", http.StatusBadRequest)
		return // exit if there was a json decode error
	}

	action := api.Action
	if action == "" {
		http.Error(w, "Missing action in request body", http.StatusBadRequest)
		return
	}

	bodyString := string(bodyBytes)

	auth := r.Header.Get("Authorization")

	claims := &Claims{}
	var err error

	if auth == "" {
		if action != "login" && action != "signup" {
			http.Error(w, "Missing Authorization Token", http.StatusUnauthorized)
			return
		}
	} else {

		claims, err = UnmarshalBearer(&auth)

		if err != nil {
			log.Printf("Error unmarshalling token: %v\n", err)
			http.Error(w, "Bad Authorization Token", http.StatusUnauthorized)
			return
		}
	}

	request := apiRequest{
		claims:       *claims,
		requestBody:  bodyString,
		response:     "",
		responseCode: http.StatusOK,
	}

	switch action {
	case "signup":
		request.signup()
	case "login":
		request.login()
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}

	if request.responseCode < 300 {
		w.Header().Add("Content-Type", "application/json")
	}

	w.WriteHeader(request.responseCode)
	w.Write([]byte(request.response))
}
