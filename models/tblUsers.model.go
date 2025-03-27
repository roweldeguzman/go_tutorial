package models

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type Users struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"firstName" gorm:"type:varchar(255)" validate:"required,min=3"`
	LastName   string `json:"lastName" gorm:"type:varchar(255)" validate:"required"`
	Email      string `json:"email" gorm:"type:varchar(255)" validate:"required,email"`
	Password   string `json:"-" gorm:"type:varchar(255)" validate:"validPassword"`
	UserStatus string `json:"-" gorm:"type:ENUM('0', '1') default '0' comment '0=For verification, 1=Verified User'"`
	DateModel
}

type UserDelete struct {
	IDS []uint
}

func (c *Users) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("email = ?", c.Email).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("User already exists")
	}

	return nil
}

func (u *Users) BeforeUpdate(tx *gorm.DB) (err error) {

	var user Users
	ctx := DB.Where("email = ?", u.Email).Find(&user)

	if ctx.RowsAffected != 0 && user.ID != u.ID {
		return errors.New("User already exists.")
	}

	return nil
}

func (c *Users) Create() error {
	return nil
	// ctx := DB.Create(&c)
	// return ctx.Error
}

func (c *Users) Update() error {
	ctx := DB.Model(&c).Where("id = ?", c.ID).Updates(Users{
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		UserStatus: c.UserStatus,
	})

	return ctx.Error
}

func (c *UserDelete) Delete() error {
	ctx := DB.Delete(&Users{}, c.IDS)
	if ctx.RowsAffected == 0 {
		return errors.New("No user deleted. User not found.")
	}
	return nil
}

func (c *Users) Get(r *http.Request) ([]Users, int64, error) {
	var users []Users
	var userCount int64

	DB.Model(&Users{}).Count(&userCount)
	ctxTotal := DB.Select("id as count").Find(&users)

	ctx := DB.Scopes(paginate(r), order(r, []string{"id", "name"})).Find(&users)

	return users, ctxTotal.RowsAffected, ctx.Error

}

func (c *Users) GetInfo() error {
	ctx := DB.Find(&c)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find user.")
	}

	return ctx.Error
}

func (c *Users) FindUser() error {

	ctx := DB.Where("email=?", c.Email).Find(&c)

	if ctx.RowsAffected == 0 {
		return errors.New("Wrong username or password.")
	}

	return ctx.Error
}
