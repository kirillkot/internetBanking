package common

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
