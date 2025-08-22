package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var savePath string

func Run(listen, savePathFlag string) {

	savePath = savePathFlag

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Static files
	e.Static("/static", "static")

	// Routes
	e.GET("/", serveUI)
	e.POST("/save", saveMarkdown)

	// Start server
	log.Printf("Server starting on port %s", listen)
	if err := e.Start(listen); err != nil {
		log.Fatal(err)
	}
}
