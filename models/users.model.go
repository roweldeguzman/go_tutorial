package models

import (
	"errors"

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
