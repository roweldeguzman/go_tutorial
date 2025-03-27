package repository

import (
	"api/models"
	"api/struct/pagination"
	"errors"
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

func (r *UserRepository) Update(user *models.Users) error {
	ctx := DB.Model(&user).Where("id = ?", user.ID).Updates(models.Users{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		UserStatus: user.UserStatus,
	})

	return ctx.Error
}

func (r *UserRepository) Delete(user *models.UserDelete) error {
	ctx := DB.Delete(&models.Users{}, user.IDS)
	if ctx.RowsAffected == 0 {
		return errors.New("No user deleted. User not found.")
	}
	return nil
}

func (r *UserRepository) Get(pageParams pagination.PageParam, sortParams pagination.SortParam) ([]models.Users, int64, error) {

	var users []models.Users
	var userCount int64

	DB.Model(&models.Users{}).Count(&userCount)

	ctx := DB.Scopes(paginate(pageParams), order(sortParams, []string{"id", "name"})).Find(&users)

	return users, userCount, ctx.Error
}

func (r *UserRepository) GetInfo(user *models.Users) error {

	ctx := DB.Find(&user)

	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find user.")
	}

	return ctx.Error
}

func (r *UserRepository) FindUser(user *models.Users) ([]models.Users, error) {
	var users []models.Users

	ctx := DB.Where("email=?", user.Email).Find(&users)

	if ctx.RowsAffected == 0 {
		return nil, errors.New("Wrong username or password.")
	}

	return users, ctx.Error
}
