package api

import (
	"backend/db"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Extract "action" from the request body (JSON)
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

func File(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		http.Error(w, "Missing 'id' parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid 'id' parameter: Must be an integer", http.StatusBadRequest)
		return
	}

	file, err := db.Connection.GetFileByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", file.Name))
	w.Header().Set("Content-Type", file.Mimetype)

	_, err = w.Write(file.Data)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
