package api

type loginResponse struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Token string `json:"token"`
}
