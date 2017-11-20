package users

import "github.com/jinzhu/gorm"

// User ...
type User struct {
	gorm.Model

	Email    string `gorm:"unique" valid:"email,required" json:"email"`
	NickName string `gorm:"unique" valid:"ascii,max=128,required" json:"nickname"`

	Password string `valid:"max=128,required" json:"password"`

	Gendor string `gorm:"gendor" valid:"" json:"gendor"`
}
