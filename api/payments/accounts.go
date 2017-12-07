package payments

import (
	"encoding/json"
	"fmt"

	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// AccountLock ...
type AccountLock struct {
	gorm.Model

	AccountID uint `gorm:"index"`
}

// Account ...
type Account struct {
	gorm.Model

	Currency string `valid:"currency,required" json:"currency"`
	Balance  int64  `valid:"required" json:"balance"`

	AddAllow  bool `json:"add_allow"`
	MoveAllow bool `json:"move_allow"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}

// IDtoIBAN ...
func IDtoIBAN(id uint) string {
	return fmt.Sprintf("BY00 00%020d", id)
}

// IBAN ...
func (a *Account) IBAN() string {
	return IDtoIBAN(a.ID)
}

// MarshalJSON ...
func (a Account) MarshalJSON() ([]byte, error) {
	type Alias Account
	return json.Marshal(struct {
		Alias
		IBAN string `json:"iban,omitempty"`
	}{
		Alias: (Alias)(a),
		IBAN:  a.IBAN(),
	})
}

// LockDB ...
func (a *Account) LockDB(tx *gorm.DB) error {
	lock, where := &AccountLock{}, &AccountLock{AccountID: a.ID}
	return tx.Set("gorm:query_option", "FOR UPDATE").
		Find(lock, where).Error
}

// AccountViewModel ...
type AccountViewModel struct{}

// NewAccountViewModel ...
func NewAccountViewModel() AccountViewModel {
	return AccountViewModel{}
}

// Name ...
func (AccountViewModel) Name() string {
	return "accounts"
}

// New ...
func (AccountViewModel) New() interface{} {
	return new(Account)
}

// NewArray ...
func (AccountViewModel) NewArray(len, cap int) interface{} {
	array := make([]Account, len, cap)
	return &array
}

// AccountView ...
type AccountView struct {
	common.ViewSet
}

// NewAccountView ...
func NewAccountView(db *gorm.DB) *AccountView {
	return &AccountView{
		ViewSet: *common.NewViewSet(db, NewAccountViewModel()),
	}
}
