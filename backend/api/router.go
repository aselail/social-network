package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// Extract the "action" from the request body (JSON)
func Router(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	auth := r.Header.Get("Authorization")

	if auth == "" {
		http.Error(w, "Missing Authorization Token", http.StatusUnauthorized)
		return
	}

	claims, err := UnmarshalBearer(&auth)

	if err != nil {
		log.Printf("Error unmarshalling token: %v\n", err)
		http.Error(w, "Bad Authorization Token", http.StatusUnauthorized)
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
