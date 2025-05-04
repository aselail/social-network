package api

type API struct {
	Action string `json:"action"`
}

// User represents a user in the system.
type User struct {
	Email    string `json:"email"`
	Password string `json:"password"` // Store hashed passwords!
}
