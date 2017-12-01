package users

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	common.Model

	UserName string `gorm:"unique" valid:"ascii,max=128,required" json:"username"`
	IsAdmin  bool   `json:"isAdmin"`

	FirstName  string `gorm:"first_name" valid:"ascii,max=128" json:"firstname"`
	SecondName string `gorm:"second_name" valid:"ascii,max=128" json:"secondname"`
	Email      string `gorm:"unique" valid:"email,required" json:"email"`
	Gendor     string `gorm:"gendor" valid:"in(m|f)" json:"gendor"`

	Password string `valid:"max=128" json:"password"`
}

// ViewModel ...
type ViewModel struct{}

// NewViewModel ...
func NewViewModel() ViewModel {
	return ViewModel{}
}

// Name ...
func (ViewModel) Name() string {
	return "users"
}

// New ...
func (ViewModel) New() interface{} {
	return new(User)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]User, len, cap)
	return &array
}

// View ...
type View struct {
	common.ViewSet
}

// NewView ...
func NewView(db *gorm.DB) *View {
	return &View{
		ViewSet: *common.NewViewSet(db, NewViewModel()),
	}
}
