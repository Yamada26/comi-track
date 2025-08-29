package main

import (
	infrastructure "comi-track/internal/infrastructure/sqlite"
	presentation "comi-track/internal/presentation/gin"
	"log"
)

func main() {
	infrastructure.InitDB("app.db")

	router := presentation.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
