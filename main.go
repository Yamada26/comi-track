package main

import (
	"comi-track/infrastructure"
	"comi-track/presentation"
	"log"
)

func main() {
	infrastructure.InitDB("app.db")

	router := presentation.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
