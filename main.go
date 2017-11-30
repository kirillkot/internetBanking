package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"internetBanking/api/common"
	"internetBanking/api/payments"
	"internetBanking/api/users"
)

var (
	port int

	logger = common.NewLogger("global")
)

func init() {
	flag.IntVar(&port, "p", 8080, "server port")
}

func migrate(db *gorm.DB) {
	tables := []interface{}{
		&users.User{},
		&payments.Account{},
	}
	for _, table := range tables {
		if err := db.CreateTable(table).Error; err != nil {
			logger.Fatalln("migrate: create table: ", err)
		}
	}
}

// NotFoundHandler - handle not found error in Angular
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {
	db := common.NewDB()
	migrate(db)
	defer db.Close()

	router := mux.NewRouter()

	users.NewView(db).RegisterRoutes(router)
	payments.NewView(db).RegisterRoutes(router)

	router.PathPrefix("/*.*").Handler(http.FileServer(http.Dir("static/")))
	router.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	logger.Infof("Stating on the %v port\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
