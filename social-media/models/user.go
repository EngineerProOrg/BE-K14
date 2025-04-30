package models

import (
	"social-media/models/sharedmodels"
	"social-media/utils"
	"time"
)

// ViewModel
type UserSignupRequestViewModel struct {
	Password string `json:"password" binding:"required"`
	sharedmodels.UserBaseViewModel
}

type UserSigninResponseViewModel struct {
	sharedmodels.UserBaseViewModel
}

type UserProfileViewModel struct {
	sharedmodels.UserBaseViewModel
}

type EditUserProfileViewModel struct {
	FirstName *string    `json:"firstName"`
	LastName  *string    `json:"lastName"`
	Birthday  *time.Time `json:"birthday"`
	Avatar    *string    `json:"avatar"`
}

// Db Model
type User struct {
	ID        int64      `gorm:"primaryKey"`
	FirstName string     `gorm:"column:first_name;size:255;not null"`
	LastName  string     `gorm:"column:last_name;size:255;not null"`
	Name      string     `gorm:"column:name;size:500;not null"`
	Birthday  time.Time  `gorm:"column:birthday;type:date"`
	Email     string     `gorm:"column:email;size:255;unique;not null"`
	Username  string     `gorm:"column:username;size:255;unique;not null"`
	Avatar    *string    `gorm:"size:500"`
	Password  string     `gorm:"column:password;size:255;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

func MapUserSignupRequestViewModelToUserDbModel(vm *UserSignupRequestViewModel) *User {
	return &User{
		FirstName: vm.FirstName,
		LastName:  vm.LastName,
		Name:      vm.FirstName + " " + vm.LastName,
		Birthday:  vm.Birthday,
		Email:     vm.Email,
		Username:  utils.GetUsernameFromEmail(vm.Email),
		Password:  vm.Password,
	}
}

func MapEditUserProfileViewModelToUserDbModel(vm *EditUserProfileViewModel) map[string]interface{} {
	updates := make(map[string]interface{})

	if vm.FirstName != nil {
		updates["first_name"] = vm.FirstName
	}
	if vm.LastName != nil {
		updates["last_name"] = vm.LastName
	}
	if vm.Birthday != nil {
		updates["birthday"] = vm.Birthday
	}
	updates["updated_at"] = time.Now()
	updates["name"] = *vm.FirstName + " " + *vm.LastName
	return updates
}

func (u *User) MapUserDbModelToUserProfileViewModel() *UserSigninResponseViewModel {
	return &UserSigninResponseViewModel{
		sharedmodels.UserBaseViewModel{
			UserId:    u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Name:      u.Name,
			Birthday:  u.Birthday,
			Email:     u.Email,
			Avatar:    u.Avatar,
		},
	}
}

func (u *User) MapUserDbModelToUserSigninResponseViewModel() *UserSigninResponseViewModel {
	return &UserSigninResponseViewModel{
		sharedmodels.UserBaseViewModel{
			UserId:    u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Name:      u.Name,
			Birthday:  u.Birthday,
			Email:     u.Email,
			Avatar:    u.Avatar,
		},
	}
}
