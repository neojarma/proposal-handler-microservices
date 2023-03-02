package models

type Reponse struct {
	Status   string    `json:"status"`
	Message  string    `json:"message"`
	Document *Document `json:"document,omitempty"`
}
