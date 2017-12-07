package payments

import (
	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// PaymentType ...
type PaymentType struct {
	common.Model

	Name string `valid:"length(4|128),required" json:"name"`

	Type      string `valid:"length(0|128),required" json:"type"`
	Commision uint   `json:"commision"`
	AccountID uint   `json:"account_id"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}

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
	return new(PaymentType)
}

// NewArray ...
func (PaymentTypeViewModel) NewArray(len, cap int) interface{} {
	array := make([]PaymentType, len, cap)
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
