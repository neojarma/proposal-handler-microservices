package models

type PayloadServices struct {
	SecretToken string
	DataRegist  *User `json:",omitempty"`
}
