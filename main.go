package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	offers "internetBanking/api/card-offers"
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
		&offers.Offer{},
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

	payments.NewView(db).RegisterRoutes(router)
	offers.NewView(db).RegisterRoutes(router)

	usersview := users.NewView(db)
	usersview.RegisterRoutes(router)

	logger.Infof("Stating server...\n")
	handler := usersview.AuthMiddleware(router)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
