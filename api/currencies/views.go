package currencies

import (
	"encoding/json"
	"internetBanking/api/models"
	"internetBanking/api/web"
	"net/http"

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
	return "currencies"
}

// New ...
func (ViewModel) New() interface{} {
	return new(models.Currency)
}

// NewArray ...
func (ViewModel) NewArray(len, cap int) interface{} {
	array := make([]models.Currency, len, cap)
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

// ConvertRequest ...
type ConvertRequest struct {
	Amount       models.Amount `json:"amount"`
	FromCurrency string        `json:"from"`
	ToCurrency   string        `json:"to"`
}

// ConvertHandler ...
func (v *View) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	request := &ConvertRequest{}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		v.Failure(w, "convert: decode: "+err.Error(), http.StatusBadRequest)
		return
	}

	result, err := models.Convert(v.DB(), request.Amount, request.FromCurrency, request.ToCurrency)
	if err != nil {
		v.Failure(w, "convert: err:"+err.Error(), http.StatusInternalServerError)
		return
	}
	model := &struct {
		Result models.Amount `json:"result"`
	}{Result: result}
	v.JSONResponse(w, model, http.StatusCreated)
}

// RegisterRoutes ...
func (v *View) RegisterRoutes(router *mux.Router, middls ...web.Middleware) {
	v.ViewSet.RegisterRoutes(router, middls...)
	router.HandleFunc("/currencies/convert/", v.ConvertHandler).Methods("GET", "POST")
}
