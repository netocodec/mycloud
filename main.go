package main

import (
	"log"
	"os"

	"./modules/config"
	_ "./modules/db"
	"./modules/webserver"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		config.DebugMode = (args[0] == "--debug")
	}

	if config.DebugMode {
		log.Println("DEBUG MODE ACTIVATED")
	}

	webserver.InitWebServer()
}
