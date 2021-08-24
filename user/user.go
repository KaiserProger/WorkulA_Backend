package user

import (
	"log"

	"gorm.io/gorm"
)

type User struct {
	UserId   int    `gorm:"primary_key" json:"user_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

var db *gorm.DB
var registeredUsers int = 0
var model *User = &User{}

func NewUser(name string, email string, password string) *User {
	registeredUsers += 1
	user := &User{
		UserId:   registeredUsers,
		Name:     name,
		Email:    email,
		Password: password,
	}
	return user
}
func GetUserByEmail(email string) *User {
	user := &User{}
	db.Model(model).Where("email = ?", email).Take(user)
	return user
}
func GetUserByID(user_id int) *User {
	user := &User{}
	db.Model(model).Where("user_id = ?", user_id).Take(user)
	return user
}
func AppendUserToDB(user *User) {
	var count int64 = 0
	db.Model(model).Where("user_id = ?", user.UserId).Count(&count)
	if count == 0 {
		db.Model(model).Create(user)
		log.Print("User " + user.Name + " created!")
	}
}
func Init(ddb *gorm.DB) {
	db = ddb
}
