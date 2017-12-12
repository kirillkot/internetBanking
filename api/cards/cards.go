package cards

import (
	"errors"
	"fmt"
	"time"

	"internetBanking/api/models"
	"internetBanking/api/payments"
	"internetBanking/api/web"

	"github.com/jinzhu/gorm"
)

// CardSimpleModel ...
type CardSimpleModel struct {
	web.BaseModel
}

// NewCardSimpleModel ...
func NewCardSimpleModel() CardSimpleModel {
	return CardSimpleModel{}
}

// Name ...
func (CardSimpleModel) Name() string {
	return "cards"
}

// New ...
func (CardSimpleModel) New() interface{} {
	return new(models.Card)
}

// NewArray ...
func (CardSimpleModel) NewArray(len, cap int) interface{} {
	array := make([]models.Card, len, cap)
	return &array
}

// CardViewModel ...
type CardViewModel struct {
	web.BaseModel
}

// NewCardViewModel ...
func NewCardViewModel() CardViewModel {
	return CardViewModel{
		BaseModel: web.NewBaseModel(NewCardSimpleModel()),
	}
}

// Create ...
func (CardViewModel) Create(db *gorm.DB, user *models.User, object interface{}) (interface{}, error) {
	card := object.(*models.Card)

	tx := db.Begin()
	if tx.Error != nil {
		return nil, errors.New("begin: " + tx.Error.Error())
	}
	defer tx.Rollback()

	offer := &models.CardOffer{}
	if err := tx.Where("id = ?", card.OfferID).Find(offer).Error; err != nil {
		return nil, errors.New("get offer: " + err.Error())
	}

	account := &models.Account{
		Currency:  card.Currency,
		Balance:   0,
		AddAllow:  true,
		MoveAllow: true,
		Detail:    fmt.Sprintf("Card Account (user %d offer %d)", user.ID, offer.ID),
	}
	if err := payments.CreateAccount(tx, account); err != nil {
		return nil, errors.New("create account: failed: " + err.Error())
	}

	now := time.Now().UTC()
	*card = models.Card{
		CardForm: card.CardForm,

		AccountID: account.ID,
		UserID:    user.ID,

		StartTime: now,
		ValidTime: now.Add(time.Duration(offer.TTL) * 30 * 24 * time.Hour),

		Type:   offer.Type,
		Status: "active",
	}
	card.SetAccount(account)
	if err := tx.Create(card).Error; err != nil {
		return nil, errors.New("create card: failed: " + err.Error())
	}

	return card, tx.Commit().Error
}

// Get ...
func (CardViewModel) Get(db *gorm.DB, user *models.User, id uint) (interface{}, error) {
	card, where := &models.Card{}, &models.Card{
		Model:  models.Model{ID: id},
		UserID: user.ID,
	}
	if err := db.Find(card, where).Error; err != nil {
		return nil, errors.New("find card: " + err.Error())
	}

	account, err := payments.GetAccountWithLock(db, card.AccountID)
	if err != nil {
		return nil, errors.New("get account with lock: " + err.Error())
	}
	card.SetAccount(account)
	return card, nil
}

// GetObjects ...
func (CardViewModel) GetObjects(db *gorm.DB, user *models.User) (interface{}, error) {
	cards, where := make([]models.Card, 0, 32), &models.Card{
		UserID: user.ID,
	}
	if err := db.Find(&cards, where).Error; err != nil {
		return nil, errors.New("get cards: " + err.Error())
	}

	for i := range cards {
		card := &cards[i]

		account, err := payments.GetAccountWithLock(db, card.AccountID)
		if err != nil {
			return nil, errors.New("get account: " + err.Error())
		}
		card.SetAccount(account)
	}
	return cards, nil
}

// CardView ...
type CardView struct {
	web.ViewSet
}

// NewCardView ...
func NewCardView(db *gorm.DB) *CardView {
	return &CardView{
		ViewSet: *web.NewViewSet(db, NewCardViewModel()),
	}
}
