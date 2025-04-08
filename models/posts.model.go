package models

import (
	"gorm.io/gorm"
)

type Posts struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Body     string `json:"firstName" gorm:"type:text" validate:"required,min=3"`
	ImageUrl string `json:"image_url" gorm:"type: text"`
	UserID   uint   `json:"-"`
	User     Users  `gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	DateModel
}

func (p *Posts) BeforeCreate(tx *gorm.DB) (err error) {
	return nil
}

func (p *Posts) BeforeUpdate(tx *gorm.DB) (err error) {
	return nil
}
