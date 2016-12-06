package models

import (
	"strings"
	"time"

	"github.com/meapex/meapex/server/db"
	"golang.org/x/crypto/bcrypt"
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
	err := db.ORM.Create(user).Error

	return user, err
}

func UpdateUser(user *User) error {
	err := db.ORM.Save(user).Error

	return err
}

func GetUserById(id interface{}) (*User, error) {
	user := new(User)
	err := db.ORM.Where("id = ?", id).First(user).Error

	return user, err
}

func GetUserByUsernameOrEmail(account string) (*User, error) {
	user := new(User)
	var err error
	if strings.Contains(account, "@") {
		err = db.ORM.Where("lower(email) = ?", strings.ToLower(account)).First(user).Error
	} else {
		err = db.ORM.Where("lower(username) = ?", strings.ToLower(account)).First(user).Error
	}

	return user, err
}

func CheckPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func generatePasswordHash(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash)
}
