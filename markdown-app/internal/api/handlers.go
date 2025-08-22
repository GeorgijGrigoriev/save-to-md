package api

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/GeorgijGrigoriev/save-to-md/internal/web"
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

// ServeFilesList serves the files list page
func serveFilesList(c echo.Context) error {
	files, err := web.FilesPage.ReadFile("files.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open files page"})
	}

	return c.HTML(http.StatusOK, string(files))
}

func serverViewUI(c echo.Context) error {
	view, err := web.ViewPage.ReadFile("view.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to open view page"})
	}

	return c.HTML(http.StatusOK, string(view))
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

	// Create directory if it doesn't exist
	err := os.MkdirAll(savePath, 0755)
	if err != nil {
		log.Printf("failed to create save directory: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create save directory"})
	}

	// Write file
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		log.Printf("save file error %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to save file"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File saved successfully", "filename": filename})
}

// ListFiles returns a list of all markdown files in the save directory
func listFiles(c echo.Context) error {
	files, err := os.ReadDir(savePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to read save directory"})
	}

	var markdownFiles []map[string]string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(strings.ToLower(file.Name()), ".md") {
			markdownFiles = append(markdownFiles, map[string]string{"name": file.Name()})
		}
	}

	return c.JSON(http.StatusOK, markdownFiles)
}

// ViewFile serves a specific markdown file for viewing
func viewFile(c echo.Context) error {
	filename := c.Param("filename")
	if filename == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Filename is required"})
	}

	filePath := filepath.Join(savePath, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "File not found"})
	}

	// Return content as plain text for viewing
	return c.String(http.StatusOK, string(content))
}

// DeleteFile handles deleting a markdown file
func deleteFile(c echo.Context) error {
	filename := c.Param("filename")
	if filename == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Filename is required"})
	}

	filePath := filepath.Join(savePath, filename)
	err := os.Remove(filePath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete file"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "File deleted successfully"})
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
