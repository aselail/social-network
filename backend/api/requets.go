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

	var request RegistrationRequest
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(ar.requestBody), &request); err != nil {
		log.Printf("Invalid request body: %v\n", err)

		ar.responseCode = http.StatusBadRequest
		ar.response = "Invalid request body"
		return
	}

	user := request.User
	user.Password = request.Password

	// Input validation (basic - expand for production)
	if strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		log.Println("Username and password are required")

		ar.responseCode = http.StatusBadRequest
		ar.response = "Username and password are required"
		return
	}

	// Check if the username already exists
	if _, err := db.Connection.FetchUserByEmail(user.Email); err == nil {
		ar.responseCode = http.StatusConflict // HTTP 409 Conflict
		ar.response = "Email already exists"
		return
	}

	if imageId, err := db.Connection.UploadImage(request.ImageFilename, request.ImageMimetype, request.ImageData); err == nil && imageId > 0 {
		user.ProfilePicture = imageId
	} else {
		log.Printf("Failed to upload Image: %v\n", err)
	}

	// Create the user
	userId, err := db.Connection.CreateUser(user)
	if err != nil {
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal Server Error"
		return
	}

	log.Printf("New user registered: %s\n", user.Email)

	c := Claims{
		Email: user.Email,
		Id:    userId,
	}

	token, err := c.GetBearer()

	if err != nil {
		log.Printf("failed to get bearer token: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	response := loginResponse{
		User:  user,
		Token: *token,
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

	var request LoginRequest
	// Use bodyString to decode the JSON
	if err := json.Unmarshal([]byte(ar.requestBody), &request); err != nil {
		log.Printf("Invalid request body: %v\n", err)

		ar.responseCode = http.StatusBadRequest
		ar.response = "Invalid request body"
		return
	}

	// Basic input validation
	if strings.TrimSpace(request.Email) == "" || strings.TrimSpace(request.Password) == "" {
		log.Println("Username and password are required")

		ar.responseCode = http.StatusBadRequest
		ar.response = "Username and password are required"
		return
	}

	// Find the user
	user, err := db.Connection.FetchUserByEmail(request.Email)

	if err != nil {
		log.Printf("Invalid credentials - user not found: %s\n", err)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	// Verify the password
	err = db.VerifyPassword(user.Password, request.Password)
	if err != nil {
		log.Printf("Password verification failed for %s: %v\n", request.Email, err)
		ar.responseCode = http.StatusUnauthorized // 401 Unauthorized
		ar.response = "Invalid credentials"
		return
	}

	log.Printf("User %s logged in\n", request.Email)
	// Successful login generating user claims and sending them

	c := Claims{
		Email: user.Email,
		Id:    user.Id,
	}

	token, err := c.GetBearer()

	if err != nil {
		log.Printf("failed to get bearer token: %v\n", err)
		ar.responseCode = http.StatusInternalServerError
		ar.response = "Internal server error"
		return
	}

	response := loginResponse{
		User:  *user,
		Token: *token,
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
