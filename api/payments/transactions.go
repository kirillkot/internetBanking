package payments

import (
	"internetBanking/api/common"
	"internetBanking/api/models"

	"github.com/jinzhu/gorm"
)

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
	return new(models.Transaction)
}

// NewArray ...
func (TransactionViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.Transaction, len, cap)
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
