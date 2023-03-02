package models

type Reponse struct {
	Status    string      `json:"status"`
	Message   string      `json:"message"`
	Buildings []*Building `json:"buildings,omitempty"`
}
