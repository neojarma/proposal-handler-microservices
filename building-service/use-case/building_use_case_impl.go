package usecase

import (
	"building-service/models"
	"building-service/repository"
)

type BuildingUseCaseImpl struct {
	Repo repository.BuildingRepo
}

func NewBuildingUseCase(repo repository.BuildingRepo) BuildingUseCase {
	return &BuildingUseCaseImpl{
		Repo: repo,
	}
}

func (useCase *BuildingUseCaseImpl) GetAll() ([]*models.Building, error) {
	return useCase.Repo.GetAll()
}

func (useCase *BuildingUseCaseImpl) FilterByStatus(status string) ([]*models.Building, error) {
	return useCase.Repo.FilterByStatus(status)
}

func (useCase *BuildingUseCaseImpl) NewBuilding(req *models.Building) error {
	return useCase.Repo.NewBuilding(req)
}

func (useCase *BuildingUseCaseImpl) UpdateBuildingStatus(req *models.Building) error {
	return useCase.Repo.UpdateBuildingStatus(req)
}
