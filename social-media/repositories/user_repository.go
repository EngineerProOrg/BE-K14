package repositories

import (
	"errors"
	"fmt"
	"social-media/models"
	"social-media/repositories/databases"
	"social-media/utils"

	"gorm.io/gorm"
)

func Signup(user *models.User) error {
	email := user.Email
	isEmailExisted, err := CheckEmailExist(email)
	if err == nil && isEmailExisted {
		return fmt.Errorf("email %s has been registered", email)
	}

	return databases.GormDb.Create(user).Error
}

func CheckEmailExist(email string) (bool, error) {
	var count int64
	err := databases.GormDb.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err // db error
	}
	return count > 0, nil // email already registered
}

func Signin(userInput *models.User) (*models.User, error) {
	var userData models.User
	err := databases.GormDb.Model(&models.User{}).Where("email = ?", userInput.Email).First(&userData).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("email or password is incorrect")
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(userInput.Password, userData.Password) {
		return nil, fmt.Errorf("email or password is incorrect")
	}
	return &userData, nil
}

func GetUserProfile(userId int64) (*models.User, error) {
	// var postEntity *models.Post // will create some unexpected exception.

	// Best practice recommend by AI.
	user := &models.User{}
	err := databases.GormDb.First(user, userId).Error
	if err != nil {
		return nil, fmt.Errorf("user does not exist")
	}

	return user, nil
}

func UpdateUserProfile(userId int64, updates map[string]interface{}) error {
	return databases.GormDb.Model(&models.User{}).Where("id = ?", userId).Updates(updates).Error
}
