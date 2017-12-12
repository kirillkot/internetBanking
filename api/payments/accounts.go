package payments

import (
	"errors"

	"internetBanking/api/models"
	"internetBanking/api/web"

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

// AccountSimpleModel ...
type AccountSimpleModel struct{}

// NewAccountSimpleModel ...
func NewAccountSimpleModel() AccountSimpleModel {
	return AccountSimpleModel{}
}

// Name ...
func (AccountSimpleModel) Name() string {
	return "accounts"
}

// New ...
func (AccountSimpleModel) New() interface{} {
	return new(models.Account)
}

// NewArray ...
func (AccountSimpleModel) NewArray(len, cap int) interface{} {
	array := make([]models.Account, len, cap)
	return &array
}

// AccountViewModel ...
type AccountViewModel struct {
	web.BaseModel
}

// NewAccountViewModel ...
func NewAccountViewModel() AccountViewModel {
	return AccountViewModel{
		BaseModel: web.NewBaseModel(NewAccountSimpleModel()),
	}
}

// Create ...
func (AccountViewModel) Create(db *gorm.DB, user *models.User, object interface{}) (interface{}, error) {
	account := object.(*models.Account)

	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, errors.New("begin: err: " + err.Error())
	}
	defer tx.Rollback()

	if err := CreateAccount(tx, account); err != nil {
		return nil, err
	}
	return account, tx.Commit().Error
}

// Get ...
func (AccountViewModel) Get(db *gorm.DB, user *models.User, id uint) (interface{}, error) {
	account, err := GetAccountWithLock(db, id)
	return account, err
}

// GetObjects ...
func (AccountViewModel) GetObjects(db *gorm.DB, user *models.User) (interface{}, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, errors.New("begin: err: " + err.Error())
	}
	defer tx.Commit()

	accounts := make([]models.Account, 0, 32)
	if err := tx.Set("gorm:query_option", "FOR UPDATE").Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}

// Delete ...
func (AccountViewModel) Delete(db *gorm.DB, user *models.User, id uint) (interface{}, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, errors.New("begin: err: " + err.Error())
	}
	defer tx.Rollback()

	account := &models.Account{}
	if err := tx.First(account, id).Error; err != nil {
		return nil, err
	}
	if err := tx.Delete(account).Error; err != nil {
		return nil, errors.New("delete account: " + err.Error())
	}
	if err := tx.Delete(&models.AccountLock{AccountID: account.ID}).Error; err != nil {
		return nil, errors.New("delete account lock: " + err.Error())
	}
	return account, tx.Commit().Error
}

// AccountView ...
type AccountView struct {
	web.ViewSet
}

// NewAccountView ...
func NewAccountView(db *gorm.DB) *AccountView {
	return &AccountView{
		ViewSet: *web.NewViewSet(db, NewAccountViewModel()),
	}
}
