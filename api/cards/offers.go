package cards

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// Offer ...
type Offer struct {
	common.Model

	Name    string `valid:"length(4|128)" json:"name"`
	Type    string `valid:"ascii,length(4|128),required" json:"type"`
	Details string `valid:"length(0|1024)" json:"details"`

	TTL      uint   `valid:"required" json:"ttlMonth"`
	Cashback uint   `json:"cashback"`
	Currency string `valid:"ascii,length(3|3)" json:"currency"`
}

// OfferViewModel ...
type OfferViewModel struct{}

// NewOfferViewModel ...
func NewOfferViewModel() OfferViewModel {
	return OfferViewModel{}
}

// Name return name for view
func (OfferViewModel) Name() string {
	return "card-offers"
}

// New ...
func (OfferViewModel) New() interface{} {
	return new(Offer)
}

// NewArray ...
func (OfferViewModel) NewArray(len, cap int) interface{} {
	array := make([]Offer, len, cap)
	return &array
}

// OfferView ...
type OfferView struct {
	common.ViewSet
}

// NewOfferView ...
func NewOfferView(db *gorm.DB) *OfferView {
	return &OfferView{
		ViewSet: *common.NewViewSet(db, NewOfferViewModel()),
	}
}
