package usecase

import (
	model "approval-service/models"
	repo "approval-service/repository"
	"time"
)

type ApprovalUseCaseImpl struct {
	Repo repo.ApprovalRepo
}

func NewApprovalUseCase(repo repo.ApprovalRepo) ApprovalUseCase {
	return &ApprovalUseCaseImpl{
		Repo: repo,
	}
}

func (useCase *ApprovalUseCaseImpl) NewDocument(req *model.Document) error {
	req.SubmissionDate = time.Now()
	return useCase.Repo.NewDocument(req)
}

func (useCase *ApprovalUseCaseImpl) UpdateProgress(req *model.Document) (int, error) {
	return useCase.Repo.UpdateProgress(req)
}

func (useCase *ApprovalUseCaseImpl) GetDocument(id string) (*model.Document, error) {
	return useCase.Repo.GetDocument(id)
}
