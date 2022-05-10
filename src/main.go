package main

import (
	"github.com/drhernandez/go-starter-project/src/server"
	"log"
)

func main() {
	if err := server.StartServer(8080); err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
