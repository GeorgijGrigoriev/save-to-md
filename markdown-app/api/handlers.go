package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/GeorgijGrigoriev/save-to-md/web"
	"github.com/labstack/echo/v4"
)

// ServeUI serves the index page
func serveUI(c echo.Context) error {
	index, err := web.IndexPage.ReadFile("index.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open index"})
	}

	return c.HTML(http.StatusOK, string(index))
}

// SaveMarkdown handles saving markdown files
func saveMarkdown(c echo.Context) error {
	// Get form data
	title := c.FormValue("title")
	content := c.FormValue("content")

	// Validate input
	if strings.TrimSpace(title) == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title is required"})
	}

	// Sanitize title to be a valid filename
	filename := sanitizeFilename(title)
	if filename == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid filename"})
	}

	filename = strings.ReplaceAll(filename, " ", "-")

	// Ensure .md extension
	if !strings.HasSuffix(filename, ".md") {
		filename += ".md"
	}

	// Create full path
	filePath := filepath.Join(savePath, filename)

	// Write file
	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Printf("save file error %v", err)

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File saved successfully", "filename": filename})
}

// sanitizeFilename removes invalid characters from filename
func sanitizeFilename(name string) string {
	// Replace invalid characters with underscore
	result := strings.Map(func(r rune) rune {
		switch r {
		case '/', '\\', ':', '*', '?', '"', '<', '>', '|':
			return '_'
		default:
			return r
		}
	}, name)

	// Remove leading/trailing whitespace and dots
	result = strings.Trim(result, " .")

	// Remove empty strings
	if result == "" {
		return ""
	}

	return result
}
