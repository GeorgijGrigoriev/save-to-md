package web

import "embed"

//go:embed index.html
var IndexPage embed.FS

//go:embed files.html
var FilesPage embed.FS

//go:embed view.html
var ViewPage embed.FS
