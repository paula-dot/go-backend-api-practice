package main

import (
	"log"

	"github.com/sikozonpc/ecom/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8080", nil)
	log.Println("Starting server on :8080")
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
