package common

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	dbURL = "postgres://postgres:postgres@postgres/db?sslmode=disable"
)

// NewDB ...
func NewDB() *gorm.DB {
	logger := NewLogger("db")

	var (
		db  *gorm.DB
		err error
	)
	for {
		db, err = gorm.Open("postgres", dbURL)
		if err == nil {
			break
		}
		logger.Fatalln("open failed:", err)
		time.Sleep(time.Second)
	}

	db.LogMode(true)
	db.SetLogger(gorm.Logger{LogWriter: logger})
	db.DB().SetMaxOpenConns(8)
	db.DB().SetMaxIdleConns(4)
	return db
}
