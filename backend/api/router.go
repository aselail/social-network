package api

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Extract the "action" from the request body (JSON)
func Router(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	auth := r.Header.Get("Authorization")

	if auth == "" {
		http.Error(w, "Missing Authorization in Header", http.StatusUnauthorized)
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

	switch action {
	case "signup":
		signupHandler(w, bodyString)
	case "login":
		loginHandler(w, bodyString)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

// signupHandler handles user registration.  Modified to receive body string.
func signupHandler(w http.ResponseWriter, bodyString string) {

	var newUser User
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(bodyString), &newUser); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Input validation (basic - expand for production)
	if strings.TrimSpace(newUser.Username) == "" || strings.TrimSpace(newUser.Password) == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Check if the username already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			http.Error(w, "Username already exists", http.StatusConflict) // HTTP 409 Conflict
			return
		}
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create a new user with the hashed password
	newUser.Password = string(hashedPassword)

	//  Store the new user (in-memory - for demonstration only)
	users = append(users, newUser)

	w.WriteHeader(http.StatusCreated) // 201 Created
	fmt.Fprintln(w, "User registered successfully")

	log.Printf("New user registered: %s", newUser.Username)
}

// loginHandler handles user login.  Modified to receive body string.
func loginHandler(w http.ResponseWriter, bodyString string) {

	var loginUser User
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(bodyString), &loginUser); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	// Basic input validation
	if strings.TrimSpace(loginUser.Username) == "" || strings.TrimSpace(loginUser.Password) == "" {
		http.Error(w, "Username and password are required", http.StatusBadRequest)
		return
	}

	// Find the user
	var foundUser *User
	for i := range users {
		if users[i].Username == loginUser.Username {
			foundUser = &users[i]
			break
		}
	}

	if foundUser == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized) // 401 Unauthorized
		return
	}

	// Verify the password
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginUser.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		log.Printf("Password verification failed for %s: %v", loginUser.Username, err)
		return
	}

	// Successful login
	fmt.Fprintln(w, "Login successful")
	log.Printf("User %s logged in", loginUser.Username)
}
