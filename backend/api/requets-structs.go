package api

type API struct {
	Action string `json:"action"`
}

// User represents a user in the system.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"` // Store hashed passwords!
}
