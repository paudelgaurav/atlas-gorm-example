package models

import (
	_ "ariga.io/atlas-provider-gorm/gormschema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"size:255;not null;unique" json:"username" form:"username" binding:"required"`
	Password       string `gorm:"not null" form:"password" json:"-"`
	NewField       string `gorm:"size:255" json:"new_field"`
	SecondNewField uint   `json:"second_new_field"`
	Role           []UserRole
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
