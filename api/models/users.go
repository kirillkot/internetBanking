package models

import (
	"encoding/json"

	"internetBanking/api/common"
)

// User ...
type User struct {
	common.Model

	Name    string `gorm:"unique" valid:"ascii,length(4|128),required" json:"username"`
	IsAdmin bool   `json:"isAdmin"`

	Password string `valid:"length(4|128)" json:"password,omitempty"`
}

// MarshalJSON ...
func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""
	type Alias User
	return json.Marshal((Alias)(u))
}
