package payments

import (
	"errors"

	"internetBanking/api/common"
	"internetBanking/api/models"

	"github.com/jinzhu/gorm"
)

// CreateAccount ...
func CreateAccount(tx *gorm.DB, account *models.Account) error {
	if err := tx.Create(account).Error; err != nil {
		return errors.New("create account: " + err.Error())
	}

	lock := &models.AccountLock{AccountID: account.ID}
	if err := tx.Create(lock).Error; err != nil {
		return errors.New("create lock: " + err.Error())
	}

	return nil
}

// GetAccountWithLock ...
func GetAccountWithLock(db *gorm.DB, id uint) (*models.Account, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, errors.New("begin: err: " + err.Error())
	}
	defer tx.Commit()

	account := &models.Account{}
	account.ID = id
	if err := account.LockDB(tx); err != nil {
		return nil, errors.New("lock: err: " + err.Error())
	}
	if err := tx.Find(account).Error; err != nil {
		return nil, errors.New("find: err: " + err.Error())
	}
	return account, nil
}

// AccountViewModel ...
type AccountViewModel struct{}

// NewAccountViewModel ...
func NewAccountViewModel() AccountViewModel {
	return AccountViewModel{}
}

// Name ...
func (AccountViewModel) Name() string {
	return "accounts"
}

// New ...
func (AccountViewModel) New() interface{} {
	return new(models.Account)
}

// NewArray ...
func (AccountViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.Account, len, cap)
	return &array
}

// AccountView ...
type AccountView struct {
	common.ViewSet
}

// NewAccountView ...
func NewAccountView(db *gorm.DB) *AccountView {
	return &AccountView{
		ViewSet: *common.NewViewSet(db, NewAccountViewModel()),
	}
}
