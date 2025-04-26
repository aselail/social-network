package api

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

// the server-side internal structure to manage and handler the request in its lifetime
type apiRequest struct {
	claims       Claims
	requestBody  string // http request body (unprocessed)
	response     string // http response
	responseCode int    // http response code (default 200)
}

// signup handles user registration.  Modified to receive body string.
func (ar *apiRequest) signup() {

	var newUser User
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(ar.requestBody), &newUser); err != nil {
		log.Printf("Invalid request body: %v\n", err)

		ar.responseCode = http.StatusBadRequest
		ar.response = "Invalid request body"
		return
	}

	// Input validation (basic - expand for production)
	if strings.TrimSpace(newUser.Username) == "" || strings.TrimSpace(newUser.Password) == "" {
		log.Println("Username and password are required")

		ar.responseCode = http.StatusBadRequest
		ar.response = "Username and password are required"
		return
	}

	// Check if the username already exists
	for _, user := range users {
		if user.Username == newUser.Username {
			log.Printf("Username already exists: %s\n", user.Username)

			ar.responseCode = http.StatusConflict // HTTP 409 Conflict
			ar.response = "Username already exists"
			return
		}
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error hashing password: %v\n", err)

		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	// Create a new user with the hashed password
	newUser.Password = string(hashedPassword)

	//  Store the new user (in-memory - for demonstration only)
	users = append(users, newUser)

	log.Printf("New user registered: %s\n", newUser.Username)

	c := Claims{
		Username: newUser.Username, // Use the new user's username
		UserID:   1,                // Replace with your actual user ID generation logic.
	}

	token, err := c.GetBearer()

	if err != nil {
		log.Printf("failed to get bearer token: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	response := loginResponse{
		Username: newUser.Username,
		UserId:   1, // TODO replace with your actual user ID logic
		Token:    *token,
	}

	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Printf("Error marshalling response: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	ar.response = string(responseJSON)
}

// login handles user login.  Modified to receive body string.
func (ar *apiRequest) login() {

	var loginUser User
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(ar.requestBody), &loginUser); err != nil {
		log.Printf("Invalid request body: %v\n", err)

		ar.responseCode = http.StatusBadRequest
		ar.response = "Invalid request body"
		return
	}

	// Basic input validation
	if strings.TrimSpace(loginUser.Username) == "" || strings.TrimSpace(loginUser.Password) == "" {
		log.Println("Username and password are required")

		ar.responseCode = http.StatusBadRequest
		ar.response = "Username and password are required"
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
		log.Printf("Invalid credentials - user not found: %s\n", loginUser.Username)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	// Verify the password
	err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginUser.Password))
	if err != nil {
		log.Printf("Password verification failed for %s: %v\n", loginUser.Username, err)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	log.Printf("User %s logged in\n", loginUser.Username)
	// Successful login generating user claims and sending them

	c := Claims{
		Username: foundUser.Username,
		UserID:   1, // TODO replace with your actual user ID logic
	}

	token, err := c.GetBearer()

	if err != nil {
		log.Printf("failed to get bearer token: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	response := loginResponse{
		Username: foundUser.Username,
		UserId:   1, // TODO replace with your actual user ID logic
		Token:    *token,
	}

	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Printf("Error marshalling response: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	ar.response = string(responseJSON)
}
