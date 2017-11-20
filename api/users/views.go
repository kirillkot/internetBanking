package users

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"internetBanking/api/common"
)

type UserView struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func (v *UserView) Get(w http.ResponseWriter, r *http.Request) {
	common.JSONResponse(w, &User{}, http.StatusOK)
}
