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
	
	// controller.Processor.Query = "Apa Mata Kuliah paling seru di semester 4"

	// Routes
	e.GET("/stimaGPT", controller.CreateQnA)
	e.GET("/stimaGPT/getResult", controller.GetResults)
	

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}




