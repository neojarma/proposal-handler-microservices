package repository

import (
	model "approval-service/models"
	"errors"

	err "approval-service/errors"

	"gorm.io/gorm"
)

type ApprovalRepoImpl struct {
	DB *gorm.DB
}

func NewApprovalRepo(db *gorm.DB) ApprovalRepo {
	return &ApprovalRepoImpl{
		DB: db,
	}
}

func (repo *ApprovalRepoImpl) NewDocument(req *model.Document) error {
	// insert data if the building status is idle
	res := repo.DB.Exec("INSERT INTO documents (document_name, building_id, submission_date, proposed_by) SELECT ?, ?, ?, ? FROM buildings WHERE buildings.id = ? AND buildings.status = 'IDLE'", req.DocumentName, req.BuildingId, req.SubmissionDate, req.ProposedBy, req.BuildingId)

	if res.RowsAffected == 0 {
		return errors.New(err.BUILDING_ALREADY_RESERVED)
	}

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (repo *ApprovalRepoImpl) UpdateProgress(req *model.Document) (int, error) {

	var buildingId int

	err := repo.DB.Model(req).Where("id = ?", req.ID).Updates(req).Select("building_id").Scan(&buildingId).Error
	if err != nil {
		return 0, err
	}

	return buildingId, nil
}

func (repo *ApprovalRepoImpl) GetDocument(id string) (*model.Document, error) {

	result := new(model.Document)

	err := repo.DB.First(result, id).Error
	if err != nil {
		return nil, err
	}

	return result, nil

}
