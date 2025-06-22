package user_dai

type UserDbModel struct {
	ID           int64  `gorm:"column:id"`
	Username     string `gorm:"column:user_name"`
	HashPassword string `gorm:"column:hash_password"`
	Email        string `gorm:"column:email"`
	DisplayName  string `gorm:"column:display_name"`
	Dob          string `gorm:"column:dob"`
	Removed      bool   `gorm:"column:removed"`
}

func (UserDbModel) TableName() string {
	return "users"
}
