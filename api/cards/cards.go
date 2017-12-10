package cards

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"internetBanking/api/models"
	"internetBanking/api/payments"
	"internetBanking/api/web"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

const (
	entry = "cards"
)

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
	return new(models.CardModel)
}

// NewArray ...
func (CardViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.CardModel, len, cap)
	return &array
}

// CardView ...
type CardView struct {
	web.ViewSet
}

// NewCardView ...
func NewCardView(db *gorm.DB) *CardView {
	return &CardView{
		ViewSet: *web.NewViewSetWithISimpleModel(db, NewCardViewModel()),
	}
}

// CreateCard ...
func CreateCard(db *gorm.DB, user *models.User, form *models.CardForm) (*models.Card, error) {
	tx := db.Begin()
	if tx.Error != nil {
		return nil, errors.New("begin: " + tx.Error.Error())
	}
	defer tx.Rollback()

	offer := &models.CardOffer{}
	if err := tx.Where("id = ?", form.OfferID).Find(offer).Error; err != nil {
		return nil, errors.New("get offer: " + err.Error())
	}

	account := &models.Account{
		Currency:  form.Currency,
		Balance:   0,
		AddAllow:  true,
		MoveAllow: true,
		Detail:    fmt.Sprintf("Card Account (user %d offer %d)", user.ID, offer.ID),
	}
	if err := payments.CreateAccount(db, account); err != nil {
		return nil, errors.New("create account: failed: " + err.Error())
	}

	now := time.Now().UTC()
	model := &models.CardModel{
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

	return &models.Card{
		CardModel: model,
		Currency:  account.Currency,
		Balance:   account.Balance,
	}, tx.Commit().Error
}

// CreateHandler ...
func (v *CardView) CreateHandler(w http.ResponseWriter, r *http.Request) {
	form := &models.CardForm{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		v.Failure(w, "create: decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(form); err != nil {
		v.Failure(w, "create: validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.UserFromRequest(r)
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
func GetCard(db *gorm.DB, user *models.User, id uint) (*models.Card, error) {
	model, where := &models.CardModel{}, &models.CardModel{
		Model:  models.Model{ID: id},
		UserID: user.ID,
	}
	if err := db.Find(model, where).Error; err != nil {
		return nil, errors.New("find model card: " + err.Error())
	}

	account, err := payments.GetAccountWithLock(db, model.AccountID)
	if err != nil {
		return nil, errors.New("get account with lock: " + err.Error())
	}

	return &models.Card{
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

	user, err := models.UserFromRequest(r)
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

// GetCards ...
func GetCards(db *gorm.DB, user *models.User) ([]models.Card, error) {
	objects, where := make([]models.CardModel, 0, 32), &models.CardModel{UserID: user.ID}
	if err := db.Find(&objects, where).Error; err != nil {
		return nil, errors.New("get card models: " + err.Error())
	}

	cards := make([]models.Card, 0, len(objects))
	for i := range objects {
		object := &objects[i]

		account, err := payments.GetAccountWithLock(db, object.AccountID)
		if err != nil {
			return nil, errors.New("get account: " + err.Error())
		}

		cards = append(cards, models.Card{
			CardModel: object,
			Balance:   account.Balance,
			Currency:  account.Currency,
		})
	}

	return cards, nil
}

// ListHandler ...
func (v *CardView) ListHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "get user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	cards, err := GetCards(v.DB(), user)
	if err != nil {
		v.Failure(w, "get cards: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.JSONResponse(w, cards, http.StatusOK)
}

// RegisterRoutes ...
func (v *CardView) RegisterRoutes(router *mux.Router) {
	prefix := "/" + entry
	router.HandleFunc(prefix+"/", v.ListHandler).Methods("GET")
	router.HandleFunc(prefix+"/", v.CreateHandler).Methods("POST")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.RetrieveHandler).Methods("GET")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.DeleteHandler).Methods("DELETE")
}
