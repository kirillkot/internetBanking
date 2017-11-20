package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"internetBanking/api/common"
)

// View ...
type View struct {
	db     *gorm.DB
	logger *logrus.Logger
}

// NewView ...
func NewView(db *gorm.DB) *View {
	return &View{
		db:     db,
		logger: common.NewLogger("users"),
	}
}

func (v *View) failure(w http.ResponseWriter, msg string, code int) {
	v.logger.Errorln(msg)
	http.Error(w, msg, code)
}

// Create ...
func (v *View) Create(w http.ResponseWriter, r *http.Request) {
	user := &User{}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		v.failure(w, "decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		v.failure(w, "validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := v.db.Create(user).Error; err != nil {
		v.failure(w, "create: "+err.Error(), http.StatusInternalServerError)
		return
	}

	common.JSONResponse(w, user, http.StatusCreated)
}

// Get ...
func (v *View) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := &User{Model: gorm.Model{ID: uint(id)}}
	switch err := v.db.First(user).Error; err {
	case nil:
		common.JSONResponse(w, user, http.StatusOK)
	case gorm.ErrRecordNotFound:
		v.failure(w, "user not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.failure(w, "get: "+err.Error(), http.StatusInternalServerError)
	}
}

func (v *View) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user := &User{Model: gorm.Model{ID: uint(id)}}
	switch err := v.db.First(user).Error; err {
	case nil:
		if e := v.db.Delete(user).Error; e != nil {
			v.failure(w, "delete: "+e.Error(), http.StatusInternalServerError)
			return
		}
		common.JSONResponse(w, user, http.StatusNoContent)
	case gorm.ErrRecordNotFound:
		v.failure(w, "user not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.failure(w, "delete: "+err.Error(), http.StatusInternalServerError)
	}
}
