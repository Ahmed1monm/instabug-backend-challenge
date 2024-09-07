package main

import (
	"fmt"

	"writer-service/db"
	"writer-service/handlers"
	"writer-service/services"
	"writer-service/services/queue"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	fmt.Fprint(e.Logger.Output(), "Hello, World!")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	db.InitDB("admin")
	db.Migrate()

	services.SetupRedis()
	queue.SetupQueue()

	e.POST("/applications/:token/chats", handlers.CreateChat)
	e.POST("/applications/:token/chats/:number/messages", handlers.CreateMessage)

	e.Logger.Fatal(e.Start(":8080"))
}
