package service

import (
	"api/models"
	"api/repository"
	"api/struct/pagination"
)

type UsersService struct {
	repo *repository.UsersRepository
}

func NewUsersService(repo *repository.UsersRepository) *UsersService {
	return &UsersService{repo}
}

func (s *UsersService) Create(user *models.Users) (*models.Users, error) {
	return s.repo.Create(user)
}

func (s *UsersService) Update(user *models.Users) error {
	return s.repo.Update(user)
}

func (s *UsersService) Delete(user *models.UserDelete) error {
	return s.repo.Delete(user)
}

func (s *UsersService) Get(pageParams pagination.PagingOptions, sortParams pagination.SortingOptions) ([]models.Users, int64, error) {
	return s.repo.Get(pageParams, sortParams)
}

func (s *UsersService) GetInfo(user *models.Users) error {
	return s.repo.GetInfo(user)
}

func (s *UsersService) SearchUser(user *models.Users) ([]models.Users, error) {
	return s.repo.SearchUser(user)
}
func (s *UsersService) FindUser(user *models.Users) (*models.Users, error) {
	return s.repo.FindUser(user)
}
