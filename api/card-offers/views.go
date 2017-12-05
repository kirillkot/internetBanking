package offers

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// Offer ...
type Offer struct {
	common.Model

	Name    string `valid:"ascii,length(4|128)" json:"name"`
	Type    string `valid:"ascii,length(4|128),required" json:"type"`
	Details string `valid:"ascii" json:"details"`

	TTL      uint   `valid:"required" json:"ttlMonth"`
	Cashback uint   `json:"cashback"`
	Currency string `valid:"ascii,length(3|3)" json:"currency"`
}

// ViewModel ...
type ViewModel struct{}

// NewViewModel ...
func NewViewModel() ViewModel {
	return ViewModel{}
}

// Name return name for view
func (ViewModel) Name() string {
	return "card-offers"
}

// New ...
func (ViewModel) New() interface{} {
	return new(Offer)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]Offer, len, cap)
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
