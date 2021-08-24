package main

import (
	"log"
	"workula/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	InitDB()
	server := echo.New()
	server.POST("/auth/signin", handlers.SignIn)
	server.POST("/auth/signup", handlers.SignUp)
	server.POST("/auth/connect", handlers.Connect)
	log.Fatal(server.Start(":1488"))
	CloseDB()
}
