package models

import (
	"errors"

	"gorm.io/gorm"
)

type Users struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	FirstName  string  `json:"firstName" gorm:"type:varchar(255)" validate:"required,min=3"`
	LastName   string  `json:"lastName" gorm:"type:varchar(255)" validate:"required"`
	Email      string  `json:"email" gorm:"type:varchar(255)" validate:"required,email"`
	Password   string  `json:"-" gorm:"type:varchar(255)" validate:"validPassword"`
	UserStatus string  `json:"-" gorm:"type:varchar(2) default '0'"`
	Post       []Posts `gorm:"foreignKey:user_id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`

	DateModel
}

type ID string

func (i ID) String() string {
	return string(i)
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
