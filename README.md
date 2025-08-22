# Markdown Editor Application

This is a web application for creating and managing markdown files through a simple web interface.

## Features

- Creating and editing markdown files
- Viewing the content of saved files in HTML format
- File management (viewing, deleting)

## Installation and Running

### Requirements
- Go 1.19 or higher
- Dependencies will be automatically downloaded on run

### Running
```bash
# Clone the repository
git clone https://github.com/GeorgijGrigoriev/save-to-md
cd markdown-app

# Start the server
go run cmd/server/main.go -listen=:8080 -savePath=/path/to/directory/to/save
```

Example startup:
```bash
go run cmd/server/main.go -listen=:8080 -savePath=/tmp/mdfiles
```

## Usage

1. Open your browser and go to `http://localhost:8080`
2. Enter the filename and content in the editor
3. Click "Save" to save the file
4. Go to the files page (`/files`) to view the list of files
5. Click on a filename to view its content in HTML format
6. Use the "Delete" button to remove files

## API Endpoints

- `GET /` - Main editor page
- `GET /files` - Files list page
- `POST /save` - Save a markdown file
- `GET /view/:filename` - View a specific file
- `DELETE /delete/:filename` - Delete a file
- `GET /api/files` - Get the list of files via API
- `GET /api/view/:filename` - Get the content of a saved file

## Project Structure

```
markdown-app/
├── cmd/
│   └── server/          # Entry point of the server
├── internal/
│   ├── api/             # API logic and handlers
└────── web/             # Web templates and static files
```