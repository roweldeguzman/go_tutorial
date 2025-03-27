package service

import (
	"api/database"
	"api/models"
	"api/repository"
	"api/struct/pagination"
	"errors"
	"fmt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {

	return &UserService{repo}
}

func (s *UserService) Create(user *models.Users) (*models.Users, error) {
	return s.repo.Create(user)
}

func (s *UserService) Update(user *models.Users) error {

	ctx := database.DB.Model(&user).Where("id = ?", user.ID).Updates(models.Users{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		UserStatus: user.UserStatus,
	})

	return ctx.Error
}

func (*UserService) Delete(user *models.UserDelete) error {
	fmt.Println(user)
	ctx := database.DB.Delete(&models.Users{}, user.IDS)
	if ctx.RowsAffected == 0 {
		return errors.New("No user deleted. User not found.")
	}
	return nil
}

func (s *UserService) Get(pageParams pagination.PageParam, sortParams pagination.SortParam) ([]models.Users, int64, error) {

	return s.repo.Get(pageParams, sortParams)

}
