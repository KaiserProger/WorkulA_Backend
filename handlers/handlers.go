package handlers

import (
	"WorkulA/db"
	"WorkulA/models/session"
	"WorkulA/models/user"
	"WorkulA/util"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Message struct {
	Message string `json:"message"`
}

func FindUserByEmail(email string) (*user.User, error) {
	u := &user.User{}
	if err := db.Connection.Model(u).Where("email = ?", email).First(u).Error; err != nil {
		return nil, err
	}
	return u, nil
}
func SignUp(c *fiber.Ctx) error {
	type UserJSON struct {
		Name     string
		Email    string
		Password string
	}
	u := &UserJSON{}
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&Message{"Error parsing user!"})
	}
	if _, err := FindUserByEmail(u.Email); err == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(&Message{"User already exists!"})
	}
	user := user.New(int(util.GenerateUserID()), u.Name, u.Email, util.GenerateHash(u.Password))
	db.Insert(user)
	s := session.New(user.UserId, u.Password)
	session.Repository.Insert(s)
	return c.Status(200).JSON(s)
}
func SignIn(c *fiber.Ctx) error {
	type LoginJSON struct {
		Email    string
		Password string
	}
	u := &LoginJSON{}
	if err := c.BodyParser(u); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&Message{"Error parsing user!"})
	}
	user, err := FindUserByEmail(u.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(&Message{"No user exists!"})
		} else {
			log.Printf("ERROR IN DB " + err.Error())
			return c.Status(fiber.StatusInternalServerError).JSON(&Message{"Something went wrong!"})
		}
	}
	if s := session.Repository.GetByUserID(user.UserId); s != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&Message{"You're already authorized!"})
	}
	h := util.GenerateHash(u.Password)
	if string(h[:]) == user.Password {
		s := session.New(user.UserId, u.Password)
		session.Repository.Insert(s)
		return c.Status(http.StatusOK).JSON(s)
	}
	return c.Status(fiber.StatusUnauthorized).JSON(&Message{"Wrong password!"})
}
