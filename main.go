package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"internetBanking/api/common"
	"internetBanking/api/payments"
	"internetBanking/api/users"
)

var (
	logger = common.NewLogger("global")
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&users.User{},
		&payments.Account{},
	).Error; err != nil {
		log.Fatalln("AutoMigrate: failed: err:", err)
	}

	admin := &users.User{
		UserName: "admin",
		IsAdmin:  true,
		Password: "admin",
	}
	if err := db.FirstOrCreate(&users.User{}, admin).Error; err != nil {
		log.Fatalln("Create admin: failed:", err)
	}
}

func main() {
	db := common.NewDB()
	migrate(db)
	defer db.Close()

	router := mux.NewRouter()

	users.NewView(db).RegisterRoutes(router)
	payments.NewView(db).RegisterRoutes(router)

	logger.Infof("Stating server...\n")
	if err := http.ListenAndServe(":80", router); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
