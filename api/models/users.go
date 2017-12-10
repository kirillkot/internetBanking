package models

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// Model ...
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

// User ...
type User struct {
	Model

	UserName  string `gorm:"unique" valid:"ascii,length(4|128),required" json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CityName  string `json:"city_name"`
	Adress    string `json:"adress"`
	IsAdmin   bool   `json:"isAdmin"`

	Password string `valid:"length(4|128)" json:"password,omitempty"`
}

// MarshalJSON ...
func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""
	type Alias User
	return json.Marshal((Alias)(u))
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
