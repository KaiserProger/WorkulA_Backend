package main

import (
	"workula/message"
	"workula/user"
	"workula/util"

	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	db = util.CreateConnection("main.db")
	user.Init(db)
	message.Init(db)
	db.AutoMigrate(&user.User{}, &user.Session{}, &message.Message{})
}
func CloseDB() {
	util.CloseConnection(db)
}
