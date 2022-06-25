package models

import (
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type TblUsers struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	FirstName  string `json:"firstName" gorm:"type:varchar(255)" validate:"required,min=3"`
	LastName   string `json:"lastName" gorm:"type:varchar(255)" validate:"required"`
	Email      string `json:"email" gorm:"type:varchar(255)" validate:"required,email"`
	Password   string `json:"-" gorm:"type:varchar(255)" validate:"required"`
	UserStatus string `json:"-" gorm:"type:ENUM('0', '1') default '0' comment '0=For verification, 1=Verified User'"`
	DateModel
}

type TblUserDelete struct {
	IDS []uint
}

func (c *TblUsers) BeforeCreate(tx *gorm.DB) (err error) {

	ctx := DB.Where("email = ?", c.Email).Find(&c)

	if ctx.RowsAffected != 0 {
		return errors.New("User already exists")
	}

	return nil
}

func (c *TblUsers) BeforeUpdate(tx *gorm.DB) (err error) {

	var user TblUsers
	ctx := DB.Where("email = ?", c.Email).Find(&user)

	if ctx.RowsAffected != 0 && user.ID != c.ID {
		return errors.New("User already exists.")
	}

	return nil
}

func (c *TblUsers) Create() error {
	ctx := DB.Create(&c)
	return ctx.Error
}

func (c *TblUsers) Update() error {
	ctx := DB.Model(&c).Where("id = ?", c.ID).Updates(TblUsers{
		FirstName:  c.FirstName,
		LastName:   c.LastName,
		Email:      c.Email,
		UserStatus: c.UserStatus,
	})

	return ctx.Error
}

func (c *TblUserDelete) Delete() error {
	ctx := DB.Delete(&TblUsers{}, c.IDS)
	if ctx.RowsAffected == 0 {
		return errors.New("No user deleted. User not found.")
	}
	return nil
}

func (c *TblUsers) Get(r *http.Request) ([]TblUsers, int64, error) {
	var sizes []TblUsers

	ctxTotal := DB.Select("id as count").Find(&sizes)

	ctx := DB.Scopes(paginate(r), order(r, []string{"id", "name"})).Find(&sizes)

	return sizes, ctxTotal.RowsAffected, ctx.Error

}

func (c *TblUsers) GetInfo() error {
	ctx := DB.Find(&c)
	if ctx.RowsAffected == 0 {
		return errors.New("Unable to find user.")
	}

	return ctx.Error
}
