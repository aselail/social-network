package api

import (
	"backend/db"
	"encoding/json"
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
	if _, err := db.Connection.FetchUserByUsername(newUser.Username); err == nil {
		ar.responseCode = http.StatusConflict // HTTP 409 Conflict
		ar.response = "Username already exists"
		return
	}
	/*	if _, err := db.Connection.FetchUserByEmail(newUser.Email); err == nil {
		ar.responseCode = http.StatusConflict // HTTP 409 Conflict
		ar.response = "Username already exists"
		return
	}*/

	// Create the user
	user := db.User{
		Username: newUser.Username,
		Password: newUser.Password,
	}

	userId, err := db.Connection.CreateUser(user)
	if err != nil {
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal Server Error"
		return
	}

	log.Printf("New user registered: %s\n", newUser.Username)

	c := Claims{
		Username: newUser.Username,
		UserId:   userId,
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
		UserId:   userId,
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
	user, err := db.Connection.FetchUserByUsername(loginUser.Username)

	if err != nil {
		log.Printf("Invalid credentials - user not found: %s\n", err)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	// Verify the password
	err = db.VerifyPassword(user.Password, loginUser.Password)
	if err != nil {
		log.Printf("Password verification failed for %s: %v\n", loginUser.Username, err)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	log.Printf("User %s logged in\n", loginUser.Username)
	// Successful login generating user claims and sending them

	c := Claims{
		Username: user.Username,
		UserId:   user.User,
	}

	token, err := c.GetBearer()

	if err != nil {
		log.Printf("failed to get bearer token: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	response := loginResponse{
		Username: user.Username,
		UserId:   user.User,
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
