package api

type loginResponse struct {
	Username string `json:"username"`
	UserId   int    `json:"userId"`
	Token    string `json:"token"`
}
