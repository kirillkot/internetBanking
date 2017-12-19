package payments

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"internetBanking/api/models"
	"internetBanking/api/web"

	"github.com/gorilla/mux"
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

func addFunds(db *gorm.DB, id uint, amount int64) (*models.Account, error) {
	tx := db.Begin()
	if err := tx.Error; err != nil {
		return nil, errors.New("begin: err: " + err.Error())
	}
	defer tx.Rollback()

	account := &models.Account{}
	account.ID = id
	if err := account.LockDB(tx); err != nil {
		return nil, errors.New("lock: err: " + err.Error())
	}
	if err := tx.Find(account).Error; err != nil {
		return nil, errors.New("find: err: " + err.Error())
	}

	if account.Balance+amount < 0 {
		return nil, errors.New("no funds")
	}

	account.Balance += amount
	if err := tx.Save(account).Error; err != nil {
		return nil, errors.New("save: err: " + err.Error())
	}

	return account, tx.Commit().Error
}

// AddFundsHandler ...
func (v *AccountView) AddFundsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "add: parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	body := &struct {
		Amount int64 `json:"amount"`
	}{}
	if err = json.NewDecoder(r.Body).Decode(body); err != nil {
		v.Failure(w, "add: decode body: "+err.Error(), http.StatusBadRequest)
		return
	}

	account, err := addFunds(v.DB(), uint(id), body.Amount)
	if err != nil {
		v.Failure(w, "add: add funds: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.JSONResponse(w, account, http.StatusCreated)
}

// CardTransactionsResponse ...
type CardTransactionsResponse struct {
	Total     int64 `json:"total"`
	TotalAdd  int64 `json:"total_add"`
	TotalMove int64 `json:"total_move"`

	Transactions []models.Transaction `json:"transactions"`
}

func accountTransactions(db *gorm.DB, account uint) ([]models.Transaction, error) {
	transactions := make([]models.Transaction, 0, 32)

	query := db.Where("account_id = ?", account).Order("time DESC").Limit(128)
	if err := query.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

// TransactionsHandler ...
func (v *AccountView) TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	transactions, err := accountTransactions(v.DB(), uint(id))
	if err != nil {
		v.Failure(w, "get: transactions: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var total, add, move int64
	for i := range transactions {
		delta := transactions[i].Delta

		total += delta
		if delta >= 0 {
			add += delta
		} else {
			move += delta
		}
	}
	model := CardTransactionsResponse{
		Total:        total,
		TotalAdd:     add,
		TotalMove:    move,
		Transactions: transactions,
	}
	v.JSONResponse(w, model, http.StatusOK)
}

// RegisterRoutes ...
func (v *AccountView) RegisterRoutes(router *mux.Router, middls ...web.Middleware) {
	v.ViewSet.RegisterRoutes(router, middls...)

	router.HandleFunc("/accounts/{id:[0-9]+}/add/",
		web.ApplyMiddl(v.AddFundsHandler, middls...)).Methods("POST")
	router.HandleFunc("/accounts/{id:[0-9]+}/transactions/",
		web.ApplyMiddl(v.TransactionsHandler, middls...)).Methods("GET")
}
