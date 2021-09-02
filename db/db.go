package db

import (
	"workula/globals"
	"workula/message"
	"workula/user"
	"workula/util"
)

func InitDB() {
	globals.DB = util.CreateConnection("main.db")
	user.Init(globals.DB)
	message.Init(globals.DB)
	globals.DB.AutoMigrate(&user.User{}, &user.Session{}, &message.Message{})
}
func CloseDB() {
	util.CloseConnection(globals.DB)
}
