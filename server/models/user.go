package models

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null;unique_index" json:"username"`
	Email     string    `gorm:"type:varchar(255);not null;unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Role      uint      `sql:"default:0"`
	IsActive  bool      `sql:"default:false" json:"is_active"`
	IsDelete  bool      `sql:"default:false" json:"is_delete"`
	Token     string    `gorm:"type:varchar(32)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateUser(username string, email string, password string) (*User, error) {
	user := &User{
		Username: username,
		Email:    email,
		Password: generatePasswordHash(password),
	}
	err := db.Create(user).Error

	return user, err
}

func UpdateUser(user *User) error {
	err := db.Save(user).Error

	return err
}

func GetUserById(id interface{}) (*User, error) {
	user := new(User)
	err := db.Where("id = ?", id).First(user).Error

	return user, err
}

func GetUserByUsernameOrEmail(account string) (*User, error) {
	user := new(User)
	var err error
	if strings.Contains(account, "@") {
		err = db.Where("lower(email) = ?", strings.ToLower(account)).First(user).Error
	} else {
		err = db.Where("lower(username) = ?", strings.ToLower(account)).First(user).Error
	}

	return user, err
}

func CheckPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func generatePasswordHash(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
		return ""
	}

	return string(passwordHash)
}
