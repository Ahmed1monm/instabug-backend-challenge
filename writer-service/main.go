package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"writer-service/models"
	"writer-service/handlers"
)


func main() {
	e := echo.New()

	fmt.Fprint(e.Logger.Output(), "Hello, World!")
	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// db connection
	models.InitDB("admin")
	models.Migrate()
	// Routes
	e.POST("/applications/:token/chats", handlers.CreateChat)
	e.POST("/applications/:token/chats/:number/messages", handlers.CreateMessage)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
