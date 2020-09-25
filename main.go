package main

import (
	"log"
	"os"

	_ "./modules/db"
	"./modules/mem"
	"./modules/webserver"
)

func main() {
	args := os.Args[1:]

	if len(args) != 0 {
		mem.DebugMode = (args[0] == "--debug")
	}

	if mem.DebugMode {
		log.Println("DEBUG MODE ACTIVATED")
	}

	webserver.InitWebServer()
}
