package main

import (
	"Backend/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)
func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/stimaGPT", controller.CreateQnA)
	e.GET("/stimaGPT/getQnA", controller.GetQnA)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}




