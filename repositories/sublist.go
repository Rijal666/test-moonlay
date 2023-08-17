package repositories

import (
	"todo-list/models"

	"gorm.io/gorm"
)

type SubListRepository interface {
	FindSubLists() ([]models.SubList, error)
	GetSubList(ID int) (models.SubList, error)
	CreateSubList(SubList models.SubList) (models.SubList, error)
	DeleteSubList(SubList models.SubList) (models.SubList, error)
	UpdateSubList(SubList models.SubList) (models.SubList, error)
	GetAllSubLists(page, pageSize int, search string, preloadSublist bool) ([]models.SubList, int, error)
}

func RepositorySubList(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSubLists() ([]models.SubList, error) {
	var sublists []models.SubList
	err := r.db.Find(&sublists).Error

	return sublists, err
}

func (r *repository) GetSubList(ID int) (models.SubList, error) {
	var sublist models.SubList
	err := r.db.First(&sublist, ID).Error

	return sublist, err
}

func (r *repository) CreateSubList(sublist models.SubList) (models.SubList, error) {
	err := r.db.Create(&sublist).Error

	return sublist, err
}

func (r *repository) DeleteSubList(sublist models.SubList) (models.SubList, error) {
	err := r.db.Delete(&sublist).Scan(&sublist).Error

	return sublist, err
}

func (r *repository) UpdateSubList(sublist models.SubList) (models.SubList, error) {
	err := r.db.Save(&sublist).Error

	return sublist, err
}

func (r *repository) GetAllSubLists(page, pageSize int, search string, preloadSublist bool) ([]models.SubList, int, error) {
	var sublists []models.SubList
	query := r.db.Model(&models.SubList{})

	if search != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	var totalSublist int64
	err := query.Count(&totalSublist).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err = query.Offset(offset).Limit(pageSize).Find(&sublists).Error
	if err != nil {
		return nil, 0, err
	}

	return sublists, int(totalSublist), nil
}
