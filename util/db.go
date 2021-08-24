package util

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CheckErrors(from string, err error) {
	if err != nil {
		log.Print(from + " ")
		log.Fatal(err)
	}
}

func CreateConnection(db_name string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{SkipDefaultTransaction: true})
	CheckErrors("CreateConnection", err)
	return db
}
func CloseConnection(db *gorm.DB) {
	ddb, err := db.DB()
	CheckErrors("CloseConnection", err)
	err = ddb.Close()
	CheckErrors("CloseConnection", err)
}
