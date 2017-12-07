package cards

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"internetBanking/api/common"
	"internetBanking/api/payments"
	"internetBanking/api/users"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	entry = "cards"
)

// CardForm ...
type CardForm struct {
	OfferID  uint   `valid:"required" json:"offer_id"`
	Name     string `valid:"length(4|128),required" json:"name"`
	Currency string `valid:"currency,required" json:"currency"`
}

// CardModel ...
type CardModel struct {
	gorm.Model

	AccountID uint `json:"account_id"`
	OfferID   uint `json:"offer_id"`
	UserID    uint `json:"user_id"`

	StartTime time.Time `json:"start_time"`
	ValidTime time.Time `json:"valid_time"`

	Name string `json:"name"`
	Type string `json:"type"`

	Status string `json:"string"`
}

// Card ...
type Card struct {
	*CardModel

	Currency string `json:"currency"`
	Balance  int64  `json:"balance"`
}

// CardViewModel ...
type CardViewModel struct{}

// NewCardViewModel ...
func NewCardViewModel() CardViewModel {
	return CardViewModel{}
}

// Name ...
func (CardViewModel) Name() string {
	return entry
}

// New ...
func (CardViewModel) New() interface{} {
	return new(CardModel)
}

// NewArray ...
func (CardViewModel) NewArray(len, cap int) interface{} {
	array := make([]CardModel, len, cap)
	return &array
}

// CardView ...
type CardView struct {
	common.ViewSet
}

// NewCardView ...
func NewCardView(db *gorm.DB) *CardView {
	return &CardView{
		ViewSet: *common.NewViewSet(db, NewCardViewModel()),
	}
}

// CreateCard ...
func CreateCard(db *gorm.DB, user *users.User, form *CardForm) (*Card, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, errors.New("begin: " + tx.Error.Error())
	}
	defer tx.Rollback()

	offer := &Offer{}
	if err := tx.Where("id = ?", form.OfferID).Find(offer).Error; err != nil {
		return nil, errors.New("get offer: " + err.Error())
	}

	account := &payments.Account{
		Currency:  form.Currency,
		Balance:   0,
		AddAllow:  true,
		MoveAllow: true,
		Detail:    fmt.Sprintf("Card Account (user %d offer %d)", user.ID, offer.ID),
	}
	if err := tx.Create(account).Error; err != nil {
		return nil, errors.New("create account: failed: " + err.Error())
	}

	now := time.Now().UTC()
	model := &CardModel{
		AccountID: account.ID,
		OfferID:   offer.ID,
		UserID:    user.ID,

		StartTime: now,
		ValidTime: now.Add(time.Duration(offer.TTL) * 30 * 24 * time.Hour),

		Name: form.Name,
		Type: offer.Type,

		Status: "active",
	}
	if err := tx.Create(model).Error; err != nil {
		return nil, errors.New("create card: failed: " + err.Error())
	}

	return &Card{
		CardModel: model,
		Currency:  account.Currency,
		Balance:   account.Balance,
	}, tx.Commit().Error
}

// CreateHandler ...
func (v *CardView) CreateHandler(w http.ResponseWriter, r *http.Request) {
	form := &CardForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		v.Failure(w, "create: decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(form); err != nil {
		v.Failure(w, "create: validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "create: get user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	card, err := CreateCard(v.DB(), user, form)
	if err != nil {
		v.Failure(w, "create: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.JSONResponse(w, card, http.StatusCreated)
}

// GetCard ...
func GetCard(db *gorm.DB, user *users.User, id uint) (*Card, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, errors.New("begin: " + tx.Error.Error())
	}
	defer tx.Rollback()

	model, where := &CardModel{}, &CardModel{Model: gorm.Model{ID: id}, UserID: user.ID}
	if err := tx.Find(model, where).Error; err != nil {
		return nil, errors.New("find model card: " + err.Error())
	}

	account := &payments.Account{Model: gorm.Model{ID: model.AccountID}}
	if err := account.LockDB(tx); err != nil {
		return nil, errors.New("account: db lock: " + err.Error())
	}

	if err := tx.Find(account).Error; err != nil {
		return nil, errors.New("account: get: " + err.Error())
	}

	return &Card{
		CardModel: model,
		Balance:   account.Balance,
		Currency:  account.Currency,
	}, nil
}

// RetrieveHandler ...
func (v *CardView) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := users.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "get user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	card, err := GetCard(v.DB(), user, uint(id))
	if err != nil {
		v.Failure(w, "get card: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.JSONResponse(w, card, http.StatusOK)
}

// ListHandler ...
func (v *CardView) ListHandler(w http.ResponseWriter, r *http.Request) {
	v.JSONResponse(w, []interface{}{}, http.StatusOK)
}

// RegisterRoutes ...
func (v *CardView) RegisterRoutes(router *mux.Router) {
	prefix := "/" + entry
	router.HandleFunc(prefix+"/", v.ListHandler).Methods("GET")
	router.HandleFunc(prefix+"/", v.CreateHandler).Methods("POST")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.RetrieveHandler).Methods("GET")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.DeleteHandler).Methods("DELETE")
}
