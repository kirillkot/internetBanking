package offers

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// Offer ...
type Offer struct {
	common.Model

	IDOffer   int
	OrderName string
	OrderType string
	Cashback  int
	Currency  int
	TTL       int
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
