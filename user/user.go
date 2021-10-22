package user

import (
	"errors"
	"log"
	"workula/objects"

	"gorm.io/gorm"
)

type User objects.User

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
	result := db.Model(model).Where("email = ?", email).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return user
}
func FindUserByName(name string) *[]User {
	users := make([]User, 0)
	result := db.Model(model).Where("name LIKE ?", name).Find(&users)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &users
}
func GetUserByID(user_id int) *User {
	user := &User{}
	db.Model(model).Where("user_id = ?", user_id).First(user)
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
