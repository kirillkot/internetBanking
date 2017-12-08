package payments

import (
	"internetBanking/api/common"
	"internetBanking/api/models"

	"github.com/jinzhu/gorm"
)

// PaymentTypeViewModel ...
type PaymentTypeViewModel struct{}

// NewPaymentTypeViewModel ...
func NewPaymentTypeViewModel() PaymentTypeViewModel {
	return PaymentTypeViewModel{}
}

// Name ...
func (PaymentTypeViewModel) Name() string {
	return "payment-types"
}

// New ...
func (PaymentTypeViewModel) New() interface{} {
	return new(models.PaymentType)
}

// NewArray ...
func (PaymentTypeViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.PaymentType, len, cap)
	return &array
}

// PaymentTypeView ...
type PaymentTypeView struct {
	common.ViewSet
}

// NewPaymentTypeView ...
func NewPaymentTypeView(db *gorm.DB) *PaymentTypeView {
	return &PaymentTypeView{
		ViewSet: *common.NewViewSet(db, NewPaymentTypeViewModel()),
	}
}
