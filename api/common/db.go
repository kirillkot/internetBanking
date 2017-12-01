package common

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	dbURL = "postgres://postgres:postgres@postgres/db?sslmode=disable"
)

// Model ...
type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

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
	db.SetLogger(gorm.Logger{logger})
	db.DB().SetMaxOpenConns(8)
	db.DB().SetMaxIdleConns(4)
	return db
}
