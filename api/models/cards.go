package models

import (
	"encoding/json"
	"time"
)

// CardOffer ...
type CardOffer struct {
	Model

	Name    string `valid:"length(4|128)" json:"name"`
	Type    string `valid:"ascii,length(4|128),required" json:"type"`
	Details string `valid:"length(0|1024)" json:"details"`

	TTL      uint   `valid:"required" json:"ttlMonth"`
	Cashback Amount `json:"cashback"`
	Currency string `valid:"ascii,length(3|3)" json:"currency"`
}

// CardForm ...
type CardForm struct {
	OfferID  uint   `valid:"required" json:"offer_id"`
	Name     string `valid:"length(4|128),required" json:"name"`
	Currency string `valid:"currency,required" json:"currency"`
}

// Card ...
type Card struct {
	Model
	CardForm

	AccountID uint `json:"account_id"`
	UserID    uint `json:"user_id"`

	StartTime time.Time `json:"start_time"`
	ValidTime time.Time `json:"valid_time"`

	Type   string `json:"type"`
	Status string `json:"status"`

	account *Account
}

// SetAccount ...
func (c *Card) SetAccount(account *Account) {
	c.account = account
}

// UnmarshalJSON ...
func (c *Card) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &c.CardForm)
}

// MarshalJSON ...
func (c Card) MarshalJSON() ([]byte, error) {
	type Alias Card
	if c.account == nil {
		return json.Marshal((Alias)(c))
	}
	return json.Marshal(struct {
		Alias
		Currency string `json:"currency"`
		Balance  Amount `json:"balance"`
	}{
		Alias:    (Alias)(c),
		Currency: c.account.Currency,
		Balance:  Amount(c.account.Balance),
	})
}
