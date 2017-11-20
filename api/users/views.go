package users

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"internetBanking/api/common"
)

type View struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewView(db *gorm.DB) *View {
	return &View{
		db:     db,
		logger: common.NewLogger("users"),
	}
}

func (v *View) Get(w http.ResponseWriter, r *http.Request) {
	common.JSONResponse(w, &User{}, http.StatusOK)
}
