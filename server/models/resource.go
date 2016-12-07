package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/meapex/meapex/server/db"
	"github.com/speps/go-hashids"
)

type Resource struct {
	ID         uint      `gorm:"primary_key" json:"id"`
	GUID       string    `gorm:"type:varchar(255);not null;unique_index" json:"guid"`
	Title      string    `gorm:"type:varchar(255);not null" json:"title"`
	Slug       string    `gorm:"type:varchar(255);not null;unique_index" json:"slug"`
	Content    string    `gorm:"type:text" json:"content,omitempty"`
	Attribute  string    `gorm:"type:text" json:"attribute,omitempty"`
	Reference  string    `sql:"type:JSON" json:"reference,omitempty"`
	Password   string    `gorm:"type:varchar(255);" json:"-"`
	IsDelete   bool      `sql:"default:false" json:"is_delete"`
	ShareCount uint      `sql:"default:0" json:"share_count"`
	ShareMode  string    `gorm:"type:varchar(255);" json:"share_mode"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (r *Resource) BeforeCreate(scope *gorm.Scope) error {
	hd := hashids.NewData()
	hd.Salt = r.Slug
	hd.MinLength = 16
	h := hashids.NewWithData(hd)
	e, _ := h.Encode([]int{0})

	scope.SetColumn("GUID", e)
	return nil
}

func (r *Resource) Create() error {
	err := db.ORM.Create(r).Error

	return err
}

func GetAllResource() (*[]Resource, error) {
	resources := []Resource{}
	err := db.ORM.Table("resources").Select("id, guid, title, slug, created_at, updated_at").Where("is_delete = ?", false).Find(&resources).Error

	return &resources, err
}

func GetResourceById(id string) (*Resource, error) {
	resource := Resource{}
	err := db.ORM.Where("guid = ?", id).First(&resource).Error

	return &resource, err
}

func GetResourceCountBySlug(slug string) (int, error) {
	count := 0
	err := db.ORM.Model(&Resource{}).Where("slug = ?", slug).Or("slug LIKE ?", slug+"-sn-%").Count(&count).Error

	return count, err
}
