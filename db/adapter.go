package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func InitDB() error {
	var err error
	Connection, err = gorm.Open(sqlite.Open("main.db"), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
func Create(v interface{}) {
	Connection.AutoMigrate(v)
}
func Insert(v interface{}) {
	Connection.Create(v)
}
func Get(v interface{}, model interface{}) {

}
