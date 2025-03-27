package repository

import (
	"api/models"
	"api/struct/pagination"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user *models.Users) (*models.Users, error) {

	ctx := DB.Create(&user)

	return user, ctx.Error
}

func (r *UserRepository) Update(user *models.Users) (*models.Users, error) {
	return user, nil
}

func (r *UserRepository) Delete(user *models.Users) (bool, error) {
	return true, nil
}

func (r *UserRepository) Get(pageParams pagination.PageParam, sortParams pagination.SortParam) ([]models.Users, int64, error) {

	var users []models.Users
	var userCount int64

	DB.Model(&models.Users{}).Count(&userCount)

	ctx := DB.Scopes(paginate(pageParams), order(sortParams, []string{"id", "name"})).Find(&users)

	return users, userCount, ctx.Error
}
