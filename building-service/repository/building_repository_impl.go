package repository

import (
	"building-service/models"

	"gorm.io/gorm"
)

type BuildingRepoImpl struct {
	DB *gorm.DB
}

func NewBuildingRepo(db *gorm.DB) BuildingRepo {
	return &BuildingRepoImpl{
		DB: db,
	}
}

func (repo *BuildingRepoImpl) GetAll() ([]*models.Building, error) {
	var result []*models.Building

	err := repo.DB.Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (repo *BuildingRepoImpl) FilterByStatus(status string) ([]*models.Building, error) {
	var result []*models.Building

	err := repo.DB.Where("status = ?", status).Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (repo *BuildingRepoImpl) NewBuilding(req *models.Building) error {
	return repo.DB.Omit("id").Create(req).Error
}

func (repo *BuildingRepoImpl) UpdateBuildingStatus(req *models.Building) error {

	return repo.DB.Model(req).Updates(req).Error

}
