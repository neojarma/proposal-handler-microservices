package repository

import model "approval-service/models"

type ApprovalRepo interface {
	NewDocument(req *model.Document) error
	UpdateProgress(req *model.Document) (int, error)
	GetDocument(id string) (*model.Document, error)
}
