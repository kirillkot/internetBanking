package models

import "github.com/asaskevich/govalidator"

var (
	currencies = []string{
		"BYN",
		"RUB",
		"USD",
		"EUR",
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

	Name     string `valid:"length(2|8)" json:"name"`
	Koef     int64  `valid:"required" json:"koef"`
	Sale     int64  `valid:"required" json:"sale"`
	Purchase int64  `valid:"required" json:"purchase"`
}
