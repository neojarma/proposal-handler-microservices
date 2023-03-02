package usecase

import "building-service/models"

type BuildingUseCase interface {
	GetAll() ([]*models.Building, error)
	FilterByStatus(status string) ([]*models.Building, error)
	NewBuilding(req *models.Building) error
	UpdateBuildingStatus(req *models.Building) error
}
