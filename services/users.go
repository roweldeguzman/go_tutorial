package services

import (
	"api/models"
)

type UserService interface {
	Create(user *models.TblUsers) (*models.TblUsers, error)
	Update(user *models.TblUsers) (*models.TblUsers, error)
	Destroy(user *models.TblUsers) (*models.TblUsers, error)
	Get() ([]models.TblUsers, error)
	GetInfo() (models.TblUsers, error)
}
