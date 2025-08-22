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
	e.GET("/files", serveFilesList)
	e.POST("/save", saveMarkdown)
	e.GET("/view/:filename", serverViewUI)
	e.GET("/api/view:filename", viewFile)
	e.DELETE("/delete/:filename", deleteFile)
	e.GET("/api/files", listFiles) // API endpoint for file listing

	// Start server
	log.Printf("Server starting on port %s", listen)
	if err := e.Start(listen); err != nil {
		log.Fatal(err)
	}
}
