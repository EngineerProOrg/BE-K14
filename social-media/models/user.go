package models

import (
	"social-media/utils"
	"time"
)

// ViewModel
type UserSignupViewModel struct {
	Name     string
	Email    string
	Password string
}

// Db Model
type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:100;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt *time.Time `gorm:"autoUpdateTime:false"`
}

func MapToUserEntity(vm UserSignupViewModel) (User, error) {
	hashPassword, err := utils.HashPassword(vm.Password)
	if err != nil {
		return User{}, err
	}
	return User{
		Name:     vm.Name,
		Email:    vm.Email,
		Password: hashPassword,
	}, nil
}
