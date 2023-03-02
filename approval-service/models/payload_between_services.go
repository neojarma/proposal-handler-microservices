package models

type PayloadServices struct {
	SecretToken  string
	DataBuilding *Building `json:",omitempty"`
	DataDocument *Document `json:",omitempty"`
}
