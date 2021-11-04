package main

import (
	"log"
	"menu/handler"
	"menu/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	storage.New("certificates/db_get.json")
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	g := e.Group("api/v1/product")
	g.GET("/:id", handler.GetProduct)
	g.GET("/", handler.GetAllProduct)
	err := e.Start(":3000")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
