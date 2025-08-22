package main

import (
	"flag"

	"github.com/GeorgijGrigoriev/save-to-md/api"
)

func main() {
	// flags for app
	listen := flag.String("listen", ":8080", "server listen port")
	saveDirFlag := flag.String("savePath", "static", "content save path")

	flag.Parse()

	api.Run(*listen, *saveDirFlag)
}
