package users

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"unique" json:"email"`
	NickName string `gorm:"unique" json:"nickname"`

	Password string `json:"-"`

	Gendor string `gorm:"gendor" json:"gendor"`
}
