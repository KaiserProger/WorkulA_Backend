package main

import (
	"log"
	"os"
	"workula/db"
	"workula/handlers"

	"github.com/labstack/echo"
)

func main() {
	db.InitDB()
	server := echo.New()
	group := server.Host("/main", echo.WrapMiddleware())
	group.
		server.POST("/auth/signin", handlers.SignIn)
	server.POST("/auth/signup", handlers.SignUp)
	server.POST("/connect", handlers.Connect)
	log.Fatal(server.Start(":" + string(os.Getenv("PORT"))))
	db.CloseDB()
}
