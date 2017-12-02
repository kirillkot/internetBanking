package users

import (
	"encoding/json"
	"internetBanking/api/common"
	"net/http"

	"github.com/asaskevich/govalidator"
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
	User  *User  `json:"user"`
	Token string `json:"token"`
}

// LoginHandler ...
func (v *View) LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := &LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		v.Failure(w, "login: decode err:"+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(request); err != nil {
		v.Failure(w, "login: validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, where := &User{}, &User{UserName: request.UserName}
	if err := v.DB().Find(user, where).Error; err != nil {
		v.Failure(w, "login: get user: "+err.Error(), http.StatusForbidden)
		return
	}
	if user.Password != request.Password {
		v.Failure(w, "login: invalid password", http.StatusForbidden)
		return
	}

	response := &LoginResponse{
		User:  user,
		Token: "temptoken",
	}

	v.JSONResponse(w, response, http.StatusCreated)
}

// RegisterRoutes ...
func (v *View) RegisterRoutes(router *mux.Router) {
	v.ViewSet.RegisterRoutes(router)
	router.HandleFunc("/users/login/", v.LoginHandler).Methods("POST")
}
