package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

var (
	currencies = []string{
		"BYN",
		"RUB",
		"USD",
		"EUR",
	}

	// DefaultCurrency ...
	DefaultCurrency = &Currency{
		Name:     "BYN",
		Koef:     1,
		Sale:     Amount(powInt64(10, precision)),
		Purchase: Amount(powInt64(10, precision)),
	}
)

func init() {
	govalidator.TagMap["currency"] = func(str string) bool {
		return govalidator.IsIn(str, currencies...)
	}
}

// Currency ...
type Currency struct {
	Model

	Name     string `gorm:"unique_index" valid:"length(2|8)" json:"name"`
	Koef     int64  `valid:"required" json:"koef"`
	Sale     Amount `valid:"required" json:"sale"`
	Purchase Amount `valid:"required" json:"purchase"`
}

func Convert(tx *gorm.DB, amount Amount, from, to string) (Amount, error) {
	cfrom := &Currency{}
	if err := tx.Where("name = ?", from).Find(cfrom).Error; err != nil {
		return 0, errors.New("currencies: " + err.Error())
	}

	cto := &Currency{}
	if err := tx.Where("name = ?", to).Find(cto).Error; err != nil {
		return 0, errors.New("find to currency: err: " + err.Error())
	}

	result := amount * cfrom.Sale / cto.Purchase
	return result, nil
}

const (
	precision = 4
)

var (
	amountKoef = decimal.New(powInt64(10, precision), 0)
)

func powInt64(x int64, n int) int64 {
	result := int64(1)
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}

// Amount ...
type Amount int64

// UnmarshalJSON ...
func (a *Amount) UnmarshalJSON(data []byte) error {
	if data[0] == '"' {
		data = data[1 : len(data)-1]
	}

	value, err := decimal.NewFromString(string(data))
	if err != nil {
		return err
	}

	*(*int64)(a) = amountKoef.Mul(value.Truncate(4)).IntPart()
	return nil
}

// MarshalJSON ...
func (a Amount) MarshalJSON() ([]byte, error) {
	return decimal.New((int64)(a), -precision).MarshalJSON()
}
