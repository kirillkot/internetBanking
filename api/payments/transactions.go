package payments

import (
	"time"

	"internetBanking/api/common"

	"github.com/jinzhu/gorm"
)

// Transaction ...
type Transaction struct {
	common.Model

	TransactionID uint      `valid:"required" json:"account_id"`
	Delta         int64     `valid:"required" json:"delta"`
	Time          time.Time `json:"time"`

	Detail string `valid:"length(0|1024)" json:"detail"`
}

// TransactionViewModel ...
type TransactionViewModel struct{}

// NewTransactionViewModel ...
func NewTransactionViewModel() TransactionViewModel {
	return TransactionViewModel{}
}

// Name ...
func (TransactionViewModel) Name() string {
	return "transactions"
}

// New ...
func (TransactionViewModel) New() interface{} {
	return new(Transaction)
}

// NewArray ...
func (TransactionViewModel) NewArray(len, cap int) interface{} {
	array := make([]Transaction, len, cap)
	return &array
}

// TransactionView ...
type TransactionView struct {
	common.ViewSet
}

// NewTransactionView ...
func NewTransactionView(db *gorm.DB) *TransactionView {
	return &TransactionView{
		ViewSet: *common.NewViewSet(db, NewTransactionViewModel()),
	}
}
