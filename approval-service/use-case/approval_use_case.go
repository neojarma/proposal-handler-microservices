package usecase

import model "approval-service/models"

type ApprovalUseCase interface {
	NewDocument(req *model.Document) error
	UpdateProgress(req *model.Document) (int, error)
	GetDocument(id string) (*model.Document, error)
}
