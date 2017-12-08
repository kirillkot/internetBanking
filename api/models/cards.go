package models

import (
	"internetBanking/api/common"
	"time"
)

// CardForm ...
type CardForm struct {
	OfferID  uint   `valid:"required" json:"offer_id"`
	Name     string `valid:"length(4|128),required" json:"name"`
	Currency string `valid:"currency,required" json:"currency"`
}

// CardModel ...
type CardModel struct {
	common.Model

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
