package service

import (
	"api/models"
	"api/repository"
	"api/struct/pagination"
	"fmt"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Create(user *models.Users) (*models.Users, error) {
	fmt.Println("this is from migration", user)

	fmt.Println(s)
	return s.repo.Create(user)
}

func (s *UserService) Update(user *models.Users) error {
	return s.repo.Update(user)
}

func (s *UserService) Delete(user *models.UserDelete) error {
	return s.repo.Delete(user)
}

func (s *UserService) Get(pageParams pagination.PagingOptions, sortParams pagination.SortingOptions) ([]models.Users, int64, error) {
	return s.repo.Get(pageParams, sortParams)
}

func (s *UserService) GetInfo(user *models.Users) error {
	return s.repo.GetInfo(user)
}

func (s *UserService) SearchUser(user *models.Users) ([]models.Users, error) {
	return s.repo.SearchUser(user)
}
func (s *UserService) FindUser(user *models.Users) (*models.Users, error) {
	return s.repo.FindUser(user)
}
