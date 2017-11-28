package common

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// IModel ...
type IModel interface {
	Name() string

	New() interface{}
	NewArray(len, cap int) interface{}
}

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
		logger: *NewLogger(model.Name()),
	}
}

// Logger ...
func (v *ViewSet) Logger() *logrus.Logger {
	return &v.logger
}

// DB ...
func (v *ViewSet) DB() *gorm.DB {
	return v.db
}

func (v *ViewSet) failure(w http.ResponseWriter, msg string, code int) {
	v.logger.Errorln(msg)
	http.Error(w, msg, code)
}

// CreateHandler ...
func (v *ViewSet) CreateHandler(w http.ResponseWriter, r *http.Request) {
	object := v.model.New()
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		v.failure(w, "create: decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	if _, err := govalidator.ValidateStruct(object); err != nil {
		v.failure(w, "create: validate: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := v.db.Create(object).Error; err != nil {
		v.failure(w, "create: db save "+err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, object, http.StatusCreated)
}

// UpdateHandler ...
func (v *ViewSet) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.failure(w, "update: parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{}, 8)
	if err := json.NewDecoder(r.Body).Decode(updates); err != nil {
		v.failure(w, "update: parse body: "+err.Error(), http.StatusBadRequest)
		return
	}

	object := v.model.New()
	switch err := v.db.Where("id = ?", id).Find(object).Error; err {
	case nil:
		if e := v.db.Model(object).Updates(updates).Error; e != nil {
			v.failure(w, "update: failed: "+e.Error(), http.StatusInternalServerError)
			return
		}
		JSONResponse(w, object, http.StatusNoContent)
	case gorm.ErrRecordNotFound:
		v.failure(w, "update: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.failure(w, "update: find object: "+err.Error(), http.StatusInternalServerError)
	}
}

// DeleteHandler ...
func (v *ViewSet) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.failure(w, "delete: parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	object := v.model.New()
	switch err := v.db.Where("id = ?", id).Find(object).Error; err {
	case nil:
		if e := v.db.Delete(object).Error; e != nil {
			v.failure(w, "delete: failed: "+e.Error(), http.StatusInternalServerError)
			return
		}
		JSONResponse(w, object, http.StatusNoContent)
	case gorm.ErrRecordNotFound:
		v.failure(w, "delete: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.failure(w, "delete: find object: "+err.Error(), http.StatusInternalServerError)
	}
}

// RetrieveHandler ...
func (v *ViewSet) RetrieveHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 32)
	if err != nil {
		v.failure(w, "parse id: "+err.Error(), http.StatusBadRequest)
		return
	}

	object := v.model.New()
	switch err := v.db.Where("id = ?", id).Find(object).Error; err {
	case nil:
		JSONResponse(w, object, http.StatusOK)
	case gorm.ErrRecordNotFound:
		v.failure(w, "get: object not found: "+strconv.FormatUint(id, 10), http.StatusNotFound)
	default:
		v.failure(w, "get: "+err.Error(), http.StatusInternalServerError)
	}
}

// ListHandler ...
func (v *ViewSet) ListHandler(w http.ResponseWriter, r *http.Request) {
	objects := v.model.NewArray(0, 32)
	switch err := v.db.Find(objects).Error; err {
	case nil:
		JSONResponse(w, objects, http.StatusOK)
	case gorm.ErrRecordNotFound:
		JSONResponse(w, []interface{}{}, http.StatusOK)
	default:
		v.failure(w, "list: "+err.Error(), http.StatusInternalServerError)
	}
}

// RegisterRoutes ...
func (v *ViewSet) RegisterRoutes(router *mux.Router) {
	prefix := "/" + v.model.Name()
	router.HandleFunc(prefix+"/", v.ListHandler).Methods("GET")
	router.HandleFunc(prefix+"/", v.CreateHandler).Methods("POST")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.RetrieveHandler).Methods("GET")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.UpdateHandler).Methods("PUT")
	router.HandleFunc(prefix+"/{id:[0-9]+}/", v.DeleteHandler).Methods("DELETE")
}

// JSONResponse ...
func JSONResponse(w http.ResponseWriter, model interface{}, code int) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(model); err != nil {
		logger.Println("JSONResponse: err:", err)
	}
}
