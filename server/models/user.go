package models

import (
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/meapex/meapex/server/db"
	"github.com/speps/go-hashids"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	GUID      string    `gorm:"type:varchar(255);not null;unique_index" json:"guid"`
	Username  string    `gorm:"type:varchar(255);not null;unique_index" json:"username"`
	Email     string    `gorm:"type:varchar(255);not null;unique_index" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"`
	Role      uint      `sql:"default:0"`
	IsActive  bool      `sql:"default:false" json:"is_active"`
	IsDelete  bool      `sql:"default:false" json:"is_delete"`
	Token     string    `gorm:"type:varchar(32)" json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(scope *gorm.Scope) error {
	hd := hashids.NewData()
	hd.Salt = u.Username
	hd.MinLength = 16
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{0})

	scope.SetColumn("GUID", e)
	return nil
}

func (u *User) Create() error {
	err := db.ORM.Create(u).Error

	return err
}

func (u *User) Update() error {
	err := db.ORM.Save(u).Error

	return err
}

func GetUserById(id string) (*User, error) {
	user := User{}
	err := db.ORM.Where("guid = ?", id).First(&user).Error

	return &user, err
}

func GetUserByUsernameOrEmail(account string) (*User, error) {
	user := User{}
	var err error
	if strings.Contains(account, "@") {
		err = db.ORM.Where("lower(email) = ?", strings.ToLower(account)).First(&user).Error
	} else {
		err = db.ORM.Where("lower(username) = ?", strings.ToLower(account)).First(&user).Error
	}

	return &user, err
}

func CheckPassword(passwordHash string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

func GeneratePasswordHash(password string) string {
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(passwordHash)
}
