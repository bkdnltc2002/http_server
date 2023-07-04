package database

import (
	"log"
	"sync"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db      *gorm.DB
	once    sync.Once
)

func GetDatabase() *gorm.DB {
	once.Do(func() {
		initConnection()
	})
	return db
}

func initConnection() {
	dsn := "root:root123@tcp(localhost:30001)/Students"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

func CloseConnection() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}
	sqlDB.Close()
}