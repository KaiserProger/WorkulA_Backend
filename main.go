package main

import (
	"log"
	"os"
	"workula/db"
	"workula/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	db.InitDB()
	server := echo.New()
	server.POST("/auth/signin", handlers.SignIn)
	server.POST("/auth/signup", handlers.SignUp)
	server.POST("/auth/connect", handlers.Connect)
	log.Fatal(server.Start(":" + string(os.Getenv("PORT"))))
	db.CloseDB()
}
