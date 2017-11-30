package users

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model

	Email    string `gorm:"unique" valid:"email,required" json:"email"`
	NickName string `gorm:"unique" valid:"ascii,max=128,required" json:"nickname"`

	Password string `valid:"max=128,required" json:"password"`

	Gendor string `gorm:"gendor" valid:"in(male|female)" json:"gendor"`
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
