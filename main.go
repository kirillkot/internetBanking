package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"internetBanking/api/common"
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
	}
	for _, table := range tables {
		if err := db.CreateTable(table).Error; err != nil {
			logger.Fatalln("migrate: create table: ", err)
		}
	}
}

func main() {
	db := common.NewDB()
	migrate(db)
	defer db.Close()

	router := mux.NewRouter()

	usersview := users.NewView(db)
	usersview.RegisterRoutes(router)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	logger.Infof("Stating on the %v port\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
