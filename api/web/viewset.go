package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"internetBanking/api/common"
	"internetBanking/api/models"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// ViewSet ...
type ViewSet struct {
	model IModel

	db     *gorm.DB
	logger logrus.Logger
}

// NewViewSet ...
func NewViewSet(db *gorm.DB, model IModel) *ViewSet {
	return &ViewSet{
		model: model,

		db:     db,
		logger: *common.NewLogger(model.Name()),
	}
}

// NewViewSetWithISimpleModel ...
func NewViewSetWithISimpleModel(db *gorm.DB, model ISimpleModel) *ViewSet {
	return NewViewSet(db, NewBaseModel(model))
}

// Logger ...
func (v *ViewSet) Logger() *logrus.Logger {
	return &v.logger
}

// DB ...
func (v *ViewSet) DB() *gorm.DB {
	return v.db
}

// Failure ...
func (v *ViewSet) Failure(w http.ResponseWriter, msg string, code int) {
	v.logger.Errorln(msg)
	http.Error(w, msg, code)
}

// JSONResponse ...
func (v *ViewSet) JSONResponse(w http.ResponseWriter, model interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		v.logger.Errorln("encode json:", err)
	}
}

// CreateHandler ...
func (v *ViewSet) CreateHandler(w http.ResponseWriter, r *http.Request) {
	object := v.model.New()
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		v.Failure(w, "create: decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(object); err != nil {
		v.Failure(w, "create: validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "create: get user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	result, err := v.model.Create(v.DB(), user, object)
	if err != nil {
		v.Failure(w, "create: model: "+err.Error(), http.StatusInternalServerError)
		return
	}

	v.JSONResponse(w, result, http.StatusCreated)
}

// UpdateHandler ...
func (v *ViewSet) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "update: parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{}, 8)
	if e := json.NewDecoder(r.Body).Decode(updates); e != nil {
		v.Failure(w, "update: parse body: "+e.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "update: get user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	result, err := v.model.Update(v.DB(), user, uint(id), updates)
	switch err {
	case nil:
		v.JSONResponse(w, result, http.StatusNoContent)
	case gorm.ErrRecordNotFound:
		v.Failure(w, "update: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.Failure(w, "update: model: "+err.Error(), http.StatusInternalServerError)
	}
}

// RetrieveHandler ...
func (v *ViewSet) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "get: user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	object, err := v.model.Get(v.DB(), user, uint(id))
	switch err {
	case nil:
		v.JSONResponse(w, object, http.StatusOK)
	case gorm.ErrRecordNotFound:
		v.Failure(w, "get: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.Failure(w, "get: model: "+err.Error(), http.StatusInternalServerError)
	}
}

// ListHandler ...
func (v *ViewSet) ListHandler(w http.ResponseWriter, r *http.Request) {
	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "get: user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	objects, err := v.model.GetObjects(v.DB(), user)
	if err != nil {
		v.Failure(w, "list: model: "+err.Error(), http.StatusInternalServerError)
		return
	}
	v.JSONResponse(w, objects, http.StatusOK)
}

// DeleteHandler ...
func (v *ViewSet) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.Failure(w, "delete: parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := models.UserFromRequest(r)
	if err != nil {
		v.Failure(w, "get: user: "+err.Error(), http.StatusUnauthorized)
		return
	}

	object, err := v.model.Delete(v.DB(), user, uint(id))
	switch err {
	case nil:
		v.JSONResponse(w, object, http.StatusOK)
	case gorm.ErrRecordNotFound:
		v.Failure(w, "delete: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.Failure(w, "delete: find object: "+err.Error(), http.StatusInternalServerError)
	}
}

// Middleware ...
type Middleware func(http.HandlerFunc) http.HandlerFunc

// ApplyMiddl ...
func ApplyMiddl(f http.HandlerFunc, middls ...Middleware) http.HandlerFunc {
	for _, middl := range middls {
		f = middl(f)
	}
	return f
}

// RegisterRoutes ...
func (v *ViewSet) RegisterRoutes(router *mux.Router, middls ...Middleware) {
	prefix := "/" + v.model.Name()
	router.HandleFunc(prefix+"/", ApplyMiddl(v.ListHandler, middls...)).Methods("GET")
	router.HandleFunc(prefix+"/", ApplyMiddl(v.CreateHandler, middls...)).Methods("POST")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", ApplyMiddl(v.RetrieveHandler, middls...)).Methods("GET")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", ApplyMiddl(v.UpdateHandler, middls...)).Methods("PUT")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", ApplyMiddl(v.DeleteHandler, middls...)).Methods("DELETE")
}
