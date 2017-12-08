package models

import (
	"encoding/json"
	"fmt"
	"internetBanking/api/common"
	"time"

	"github.com/jinzhu/gorm"
)

// AccountLock ...
type AccountLock struct {
	common.Model

	AccountID uint `gorm:"index"`
}

// Account ...
type Account struct {
	common.Model

	Currency string `valid:"currency,required" json:"currency"`
	Balance  int64  `json:"balance"`

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

// PaymentType ...
type PaymentType struct {
	common.Model

	Name string `valid:"length(4|128),required" json:"name"`

	Type      string `valid:"length(0|128),required" json:"type"`
	Commision uint   `json:"commision"`
	AccountID uint   `json:"account_id"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}

// Transaction ...
type Transaction struct {
	common.Model

	TransactionID uint      `valid:"required" json:"account_id"`
	Delta         int64     `valid:"required" json:"delta"`
	Time          time.Time `json:"time"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}
