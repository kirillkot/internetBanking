package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"internetBanking/api/cards"
	"internetBanking/api/common"
	"internetBanking/api/models"
	"internetBanking/api/payments"
	"internetBanking/api/users"
)

var (
	logger = common.NewLogger("global")
)

func migrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.AccountLock{},
		&models.Transaction{},
		&models.CardOffer{},
		&models.Card{},
		&models.PaymentType{},
		&models.Payment{},
	).Error; err != nil {
		log.Fatalln("AutoMigrate: failed: err:", err)
	}

	admin := &models.User{
		Name:     "admin",
		IsAdmin:  true,
		Password: "admin",
	}
	if err := db.FirstOrCreate(&models.User{}, admin).Error; err != nil {
		log.Fatalln("Create admin: failed:", err)
	}
	if err := db.FirstOrCreate(models.BankAccount, models.BankAccount).Error; err != nil {
		log.Fatalln("Create bank account:", err)
	}
}

func main() {
	db := common.NewDB()
	migrate(db)
	defer db.Close()

	router := mux.NewRouter()

	payments.NewAccountView(db).RegisterRoutes(router)
	payments.NewTransactionView(db).RegisterRoutes(router)
	payments.NewPaymentTypeView(db).RegisterRoutes(router)
	payments.NewPaymentView(db).RegisterRoutes(router)
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
