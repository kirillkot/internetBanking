package users

import (
	"encoding/json"
	"internetBanking/api/common"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	common.Model

	UserName string `gorm:"unique" valid:"ascii,length(4|128),required" json:"username"`
	IsAdmin  bool   `json:"isAdmin"`

	Password string `valid:"length(4|128)" json:"password,omitempty"`
}

// MarshalJSON ...
func (u User) MarshalJSON() ([]byte, error) {
	u.Password = ""
	type Alias User
	return json.Marshal((Alias)(u))
}

// ViewModel ...
type ViewModel struct{}

// NewViewModel ...
func NewViewModel() ViewModel {
	return ViewModel{}
}

// Name ...
func (ViewModel) Name() string {
	return "users"
}

// New ...
func (ViewModel) New() interface{} {
	return new(User)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]User, len, cap)
	return &array
}

// View ...
type View struct {
	common.ViewSet
}

// NewView ...
func NewView(db *gorm.DB) *View {
	return &View{
		ViewSet: *common.NewViewSet(db, NewViewModel()),
	}
}

// LoginRequest ...
type LoginRequest struct {
	UserName string `valid:"required" json:"username"`
	Password string `valid:"required" json:"password"`
}

// LoginResponse ...
type LoginResponse struct {
	IsAdmin bool `json:"is_admin"`
}

// LoginHandler ...
func (v *View) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user, err := UserFromRequest(r)
	if err != nil {
		v.Failure(w, "Get user failed (are you use auth middle?): "+err.Error(), http.StatusUnauthorized)
		return
	}

	response := &LoginResponse{
		IsAdmin: user.IsAdmin,
	}
	v.JSONResponse(w, response, http.StatusCreated)
}

// RegisterRoutes ...
func (v *View) RegisterRoutes(router *mux.Router) {
	v.ViewSet.RegisterRoutes(router)
	router.HandleFunc("/login/", v.LoginHandler).Methods("GET", "POST")
}
