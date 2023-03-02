package models

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
