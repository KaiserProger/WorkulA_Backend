package main

import (
	"WorkulA/db"
	"WorkulA/handlers"
	"WorkulA/models/session"
	"WorkulA/models/user"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}
	db.Create(&user.User{})
	db.Create(&session.Session{})
	session.Repository.Init()
	app := fiber.New()
	app.Post("/auth/signup", handlers.SignUp)
	app.Post("/auth/signin", handlers.SignIn)
	app.Post("auth/signout", nil)
	log.Fatal(app.Listen(":8800" /*fmt.Sprintf(":%s", os.Getenv("PORT"))*/))
}
