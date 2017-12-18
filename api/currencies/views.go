package currencies

import (
	"internetBanking/api/models"
	"internetBanking/api/web"

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
