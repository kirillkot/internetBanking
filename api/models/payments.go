package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)

var (
	// BankAccount ...
	BankAccount = &Account{
		AddAllow:  true,
		MoveAllow: false,
		Currency:  "BYN",
		Detail:    "account for bank commisions",
	}
)

// AccountLock ...
type AccountLock struct {
	Model

	AccountID uint `gorm:"index,unique"`
}

// Account ...
type Account struct {
	Model

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

const (
	CommisionKoef = 100
)

// PaymentType ...
type PaymentType struct {
	Model

	Name string `valid:"length(4|128),required" json:"name"`

	Type      string `valid:"length(0|128),required" json:"type"`
	Commision uint   `json:"commision"`
	AccountID uint   `json:"account_id"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}

// PaymentForm ...
type PaymentForm struct {
	TypeID uint   `valid:"required" json:"payment_type_id"`
	Name   string `valid:"length(4|128),required" json:"name"`

	FromAccountID uint `valid:"required" json:"from_account_id"`

	Currency string `valid:"currency,required" json:"currency"`
	Amount   int64  `valid:"required" json:"amount"`
}

// Payment ...
type Payment struct {
	Model
	PaymentForm

	Type      string `json:"type"`
	Commision int64  `json:"commision"`

	UserID uint   `json:"-"`
	From   string `json:"from"`
	To     string `json:"to"`
}

// UnmarshalJSON ...
func (p *Payment) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &p.PaymentForm)
}

// Transaction ...
type Transaction struct {
	Model

	AccountID uint      `valid:"required" json:"account_id"`
	Delta     int64     `valid:"required" json:"delta"`
	Time      time.Time `json:"time"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}
