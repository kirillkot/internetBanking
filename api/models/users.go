package models

import (
	"context"
	"crypto"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sec51/twofactor"
)

// Model ...
type Model struct {
	ID uint `gorm:"primary_key" json:"id"`
}

const (
	issuer = "G&K"
)

// User ...
type User struct {
	Model

	Name      string `gorm:"unique" valid:"ascii,length(4|128),required" json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CityName  string `json:"city_name"`
	Adress    string `json:"adress"`
	IsAdmin   bool   `json:"isAdmin"`

	Password  string `valid:"length(4|128)" json:"password,omitempty"`
	TwoFactor []byte `valid:"-" json:"-"`
}

// GenerateTwoFactor ...
func (u *User) GenerateTwoFactor() error {
	otp, err := twofactor.NewTOTP(u.Name, issuer, crypto.SHA1, 6)
	if err != nil {
		return err
	}
	data, err := otp.ToBytes()
	if err != nil {
		return err
	}
	u.TwoFactor = data
	return nil
}

// ValidateTwoFactor ...
func (u *User) ValidateTwoFactor(code string) error {
	otp, err := twofactor.TOTPFromBytes(u.TwoFactor, issuer)
	if err != nil {
		return err
	}
	return otp.Validate(code)
}

// QR ...
func (u *User) QR() ([]byte, error) {
	otp, err := twofactor.TOTPFromBytes(u.TwoFactor, issuer)
	if err != nil {
		return nil, err
	}
	return otp.QR()
}

// MarshalJSON ...
func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""

	qr, _ := u.QR()
	type Alias User
	return json.Marshal(struct {
		Alias
		QR []byte `json:"qr,string"`
	}{
		Alias: (Alias)(u),
		QR:    qr,
	})
}

const (
	userkey = "user"
)

// SetUserToRequest ...
func SetUserToRequest(req *http.Request, user *User) *http.Request {
	ctx := context.WithValue(req.Context(), userkey, user)
	return req.WithContext(ctx)
}

// UserFromRequest ...
func UserFromRequest(req *http.Request) (*User, error) {
	user, ok := req.Context().Value(userkey).(*User)
	if !ok {
		return nil, errors.New("user is not set")
	}
	return user, nil
}
