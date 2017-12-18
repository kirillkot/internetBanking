package users

import (
	"encoding/json"
	"errors"
	"net/http"

	"internetBanking/api/models"
	"internetBanking/api/web"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

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
	return new(models.User)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.User, len, cap)
	return &array
}

// View ...
type View struct {
	web.ViewSet
}

// NewView ...
func NewView(db *gorm.DB) *View {
	return &View{
		ViewSet: *web.NewViewSetWithISimpleModel(db, NewViewModel()),
	}
}

// LoginCreds ...
type LoginCreds struct {
	UserName string `valid:"required" json:"username"`
	Password string `valid:"required" json:"password"`
}

func checkLoginCreds(db *gorm.DB, creds *LoginCreds) (*models.User, error) {
	user := &models.User{}
	if err := db.Where("name = ?", creds.UserName).Find(user).Error; err != nil {
		return nil, errors.New("get user: " + err.Error())
	}

	if user.Password != creds.Password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

// LoginHandler ...
func (v *View) LoginHandler(w http.ResponseWriter, r *http.Request) {
	creds := &LoginCreds{}
	if err := json.NewDecoder(r.Body).Decode(creds); err != nil {
		v.Failure(w, "parse creds err: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := checkLoginCreds(v.DB(), creds)
	if err != nil {
		v.Failure(w, "check creds: "+err.Error(), http.StatusForbidden)
		return
	}

	setAuthCookie(w, user)
	v.JSONResponse(w, user, http.StatusCreated)
}

// MeHandler ...
func (v *View) MeHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "Get user failed (are you use auth middle?): "+err.Error(), http.StatusUnauthorized)
		return
	}

	v.JSONResponse(w, user, http.StatusOK)
}

// RegisterRoutes ...
func (v *View) RegisterRoutes(router *mux.Router, middls ...web.Middleware) {
	v.ViewSet.RegisterRoutes(router, middls...)
	router.HandleFunc("/me/", web.ApplyMiddl(v.MeHandler, middls...)).Methods("GET")

	router.HandleFunc("/login/", v.LoginHandler).Methods("GET", "POST")
}
