package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"internetBanking/api/cards"
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
		&payments.AccountLock{},
		&payments.Transaction{},
		&cards.Offer{},
		&cards.CardModel{},
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

	payments.NewAccountView(db).RegisterRoutes(router)
	payments.NewTransactionView(db).RegisterRoutes(router)
	cards.NewOfferView(db).RegisterRoutes(router)
	cards.NewCardView(db).RegisterRoutes(router)

	usersview := users.NewView(db)
	usersview.RegisterRoutes(router)

	logger.Infof("Stating server...\n")
	handler := usersview.AuthMiddleware(router)
	if err := http.ListenAndServe(":80", handler); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
