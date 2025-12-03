package user

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string `gorm:"size:100;not null;uniqueIndex" json:"username"`
	Email    string `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Password string `gorm:"size:100;not null;" json:"-"`
}
