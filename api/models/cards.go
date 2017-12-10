package models

import (
	"time"
)

// CardOffer ...
type CardOffer struct {
	Model

	Name    string `valid:"length(4|128)" json:"name"`
	Type    string `valid:"ascii,length(4|128),required" json:"type"`
	Details string `valid:"length(0|1024)" json:"details"`

	TTL      uint   `valid:"required" json:"ttlMonth"`
	Cashback uint   `json:"cashback"`
	Currency string `valid:"ascii,length(3|3)" json:"currency"`
}

// CardForm ...
type CardForm struct {
	OfferID  uint   `valid:"required" json:"offer_id"`
	Name     string `valid:"length(4|128),required" json:"name"`
	Currency string `valid:"currency,required" json:"currency"`
}

// CardModel ...
type CardModel struct {
	Model

	AccountID uint `json:"account_id"`
	OfferID   uint `json:"offer_id"`
	UserID    uint `json:"user_id"`

	StartTime time.Time `json:"start_time"`
	ValidTime time.Time `json:"valid_time"`

	Name string `json:"name"`
	Type string `json:"type"`

	Status string `json:"string"`
}

// Card ...
type Card struct {
	*CardModel

	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}
