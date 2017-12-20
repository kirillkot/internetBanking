package payments

import (
	"errors"
	"fmt"
	"time"

	"internetBanking/api/models"
	"internetBanking/api/web"

	"github.com/jinzhu/gorm"
)

// PaymentSimpleModel ...
type PaymentSimpleModel struct {
	web.BaseModel
}

// NewPaymentSimpleModel ...
func NewPaymentSimpleModel() PaymentSimpleModel {
	return PaymentSimpleModel{}
}

// Name ...
func (PaymentSimpleModel) Name() string {
	return "payments"
}

// New ...
func (PaymentSimpleModel) New() interface{} {
	return new(models.Payment)
}

// NewArray ...
func (PaymentSimpleModel) NewArray(len, cap int) interface{} {
	array := make([]models.Payment, len, cap)
	return &array
}

// PaymentViewModel ...
type PaymentViewModel struct {
	web.BaseModel
}

// NewPaymentViewModel ...
func NewPaymentViewModel() PaymentViewModel {
	return PaymentViewModel{
		BaseModel: web.NewBaseModel(NewPaymentSimpleModel()),
	}
}

func moveFunds(tx *gorm.DB, fromID, toID uint, amount models.Amount, currency string) error {
	locks := make([]models.AccountLock, 0, 2)
	query := tx.Set("gorm:query_option", "FOR UPDATE").
		Where("account_id in (?)", []uint{fromID, toID})
	if err := query.Find(&locks).Error; err != nil {
		return errors.New("lock records: " + err.Error())
	}

	from, to := &models.Account{}, &models.Account{}
	if err := tx.Find(from, fromID).Find(to, toID).Error; err != nil {
		return errors.New("accounts: " + err.Error())
	}

	if from.Balance < amount {
		return errors.New("account: no funds")
	}
	from.Balance -= amount

	converted, err := models.Convert(tx, amount, from.Currency, to.Currency)
	if err != nil {
		return errors.New("convert: " + err.Error())
	}
	to.Balance += converted

	if err := tx.Save(from).Save(to).Error; err != nil {
		return errors.New("save: " + err.Error())
	}

	now := time.Now().UTC()
	if err := tx.Save(&models.Transaction{
		AccountID: from.ID,
		Delta:     -amount,
		Time:      now,
		Detail:    fmt.Sprintf("Move to %s (%s)", to.IBAN(), to.Detail),
	}).Save(&models.Transaction{
		AccountID: to.ID,
		Delta:     converted,
		Time:      now,
		Detail:    fmt.Sprintf("Move from %s (%s)", from.IBAN(), from.Detail),
	}).Error; err != nil {
		return errors.New("save transactions: " + err.Error())
	}
	return nil
}

// Create ...
func (PaymentViewModel) Create(db *gorm.DB, user *models.User, object interface{}) (interface{}, error) {
	payment := object.(*models.Payment)
	form := &payment.PaymentForm

	tx := db.Begin()
	if tx.Error != nil {
		return nil, errors.New("begin: " + tx.Error.Error())
	}
	defer tx.Rollback()

	paymentType := &models.PaymentType{}
	if err := tx.Find(paymentType, form.TypeID).Error; err != nil {
		return nil, errors.New("payment type: " + err.Error())
	}

	if err := moveFunds(tx, form.FromAccountID, paymentType.AccountID,
		form.Amount, form.Currency); err != nil {
		return nil, err
	}

	commision := (form.Amount * paymentType.Commision) / models.CommisionKoef
	if commision != 0 {
		if err := moveFunds(tx, form.FromAccountID, models.BankAccount.ID,
			commision, form.Currency); err != nil {
			return nil, err
		}
	}

	*payment = models.Payment{
		PaymentForm: payment.PaymentForm,
		Type:        paymentType.Name,
		Commision:   commision,

		UserID: user.ID,
		From:   models.IDtoIBAN(form.FromAccountID),
		To:     models.IDtoIBAN(paymentType.AccountID),
	}
	if err := tx.Save(payment).Error; err != nil {
		return nil, errors.New("save payment: " + err.Error())
	}
	return payment, tx.Commit().Error
}

// GetObjects ...
func (PaymentViewModel) GetObjects(db *gorm.DB, user *models.User) (interface{}, error) {
	query := db
	if !user.IsAdmin {
		query = query.Where("user_id = ?", user.ID)
	}

	objects := make([]models.Payment, 0, 32)
	if err := query.Find(&objects).Error; err != nil {
		return nil, err
	}
	return objects, nil
}

// Delete ...
func (PaymentViewModel) Delete(db *gorm.DB, user *models.User, id uint) (object interface{}, err error) {
	return nil, errors.New("not implemented")
}

// PaymentView ...
type PaymentView struct {
	web.ViewSet
}

// NewPaymentView ...
func NewPaymentView(db *gorm.DB) *PaymentView {
	return &PaymentView{
		ViewSet: *web.NewViewSet(db, NewPaymentViewModel()),
	}
}
