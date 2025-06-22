package user_dai

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"ep.k14/newsfeed/internal/service/model"
)

type UserDAI struct {
	db *gorm.DB
}

type UserDbConfig struct {
	Username     string
	Password     string
	Host         string
	Port         int
	DatabaseName string
}

func New(conf *UserDbConfig) (*UserDAI, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %s", err)
	}

	return &UserDAI{
		db: db,
	}, nil
}

func (d *UserDAI) Signup(ctx context.Context, user *model.User) (*model.User, error) {
	dbUser := &UserDbModel{
		Username:     user.Username,
		HashPassword: hash(user.Password),
		Email:        user.Email,
		DisplayName:  user.DisplayName,
		Dob:          user.Dob,
		Removed:      false,
	}

	result := d.db.Create(dbUser)
	if err := result.Error; err != nil {
		fmt.Printf("error when creating user in db: %s\n", err)
		return nil, fmt.Errorf("failed to create user: %s", err)
	}
	if result.RowsAffected == 0 {
		fmt.Println("no created record")
		return nil, fmt.Errorf("no created record")
	}

	user.ID = dbUser.ID
	return user, nil
}

func hash(raw string) string {
	// TODO: hash
	return raw
}
