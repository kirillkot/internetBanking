package cards

import (
	"internetBanking/api/common"
	"internetBanking/api/models"

	"github.com/jinzhu/gorm"
)

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
	return new(models.CardOffer)
}

// NewArray ...
func (OfferViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.CardOffer, len, cap)
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
