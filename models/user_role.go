package models

import (
	"gorm.io/gorm"
)

type UserRole struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	User   User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Role   string `json:"role"`
}
