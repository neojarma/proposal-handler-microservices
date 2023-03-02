package models

import "time"

type Document struct {
	ID             string    `json:"id,omitempty"`
	DocumentName   string    `json:"documentName"`
	BuildingId     int       `json:"buildingId"`
	Status         string    `json:"status"`
	SubmissionDate time.Time `json:"submissionDate"`
	VerifiedBy     string    `json:"verifiedBy,omitempty"`
	ProposedBy     string    `json:"proposedBy"`
}
