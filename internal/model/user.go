package model

import (
	"gorm.io/gorm"
	"short-url/tools"
)

type User struct {
	Model
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100); not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Password = tools.Encrypt(u.Password)
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.Password = tools.Encrypt(u.Password)
	return nil
}
