package payments

import (
	"encoding/json"
	"fmt"

	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// Account ...
type Account struct {
	gorm.Model

	IBAN     string `gorm:"-" json:"pan,omitempty"`
	Currency string `valid:"currency,required" json:"currency"`
	Cents    int64  `valid:"required" json:"cents"`
}

// IDtoIBAN ...
func IDtoIBAN(id uint) string {
	return fmt.Sprintf("BY00 00%020d", id)
}

// MarshalJSON ...
func (a Account) MarshalJSON() ([]byte, error) {
	a.IBAN = IDtoIBAN(a.ID)

	type Alias Account
	return json.Marshal((Alias)(a))
}

// ViewModel ...
type ViewModel struct{}

// NewViewModel ...
func NewViewModel() ViewModel {
	return ViewModel{}
}

// Name ...
func (ViewModel) Name() string {
	return "accounts"
}

// New ...
func (ViewModel) New() interface{} {
	return new(Account)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]Account, len, cap)
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
