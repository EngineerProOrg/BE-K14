package repositories

import (
	"errors"
	"fmt"
	"social-media/models"
	"social-media/repositories/databases"
	"social-media/utils"

	"gorm.io/gorm"
)

func Signup(user *models.User) (*models.User, error) {
	email := user.Email
	isEmailExisted, err := CheckEmailExist(email)
	if err == nil && isEmailExisted {
		return nil, fmt.Errorf("email %s has been registered", email)
	}

	if err := databases.GormDb.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
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
			return nil, utils.ErrInvalidLogin
		}
		return nil, err
	}
	if !utils.CheckPasswordHash(userInput.Password, userData.Password) {
		return nil, utils.ErrInvalidLogin
	}
	return &userData, nil
}

func GetUserProfile(userId int64) (*models.User, error) {
	// Best practice recommend by AI.
	user := &models.User{}
	err := databases.GormDb.First(user, userId).Error
	if err != nil {
		return nil, utils.ErrUserDoesNotExist
	}

	return user, nil
}

func GetUserProfileByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := databases.GormDb.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, utils.ErrUserDoesNotExist
	}

	return user, nil
}

func UpdateUserProfile(userId int64, updates map[string]interface{}) error {
	return databases.GormDb.Model(&models.User{}).Where("id = ?", userId).Updates(updates).Error
}

func CheckUserExist(userId int64) (*models.User, error) {
	user := &models.User{}
	err := databases.GormDb.First(&user, userId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.ErrUserDoesNotExist
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
