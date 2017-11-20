package main

import (
	"flag"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func main() {
	db := common.NewDB()

	router := mux.NewRouter()

	usersview := users.NewView(db)
	router.HandleFunc("/users/{id:[0-9]+}", usersview.Get).Methods("GET")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	logger.Infof("Stating on the %v port\n", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), router); err != nil {
		logger.Fatalln("Listen and Serve: err:", err)
	}
}
