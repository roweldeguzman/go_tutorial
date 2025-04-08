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
	UserStatus string `json:"-" gorm:"type:varchar(2) default '0'"`
	Post       []Posts

	DateModel
}

// type userStatus uint

// const (
// 	forConfirmation userStatus = 0 // Represents a user pending confirmation
// 	active          userStatus = 1 // Represents a user user is active
// 	inActive        userStatus = 2 // Represents a user user is block
// )

type UserDelete struct {
	IDS []uint
}

func (c *Users) BeforeCreate(tx *gorm.DB) (err error) {
	// if c.UserStatus == 0 {
	// 	c.UserStatus = forConfirmation
	// }

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
