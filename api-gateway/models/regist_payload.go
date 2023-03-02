package models

type RegistPayload struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
