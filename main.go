package main

import (
	"HeadHunter/internal/delivery"
	"log"
)

func main() {
	// running server
	err := delivery.RunGRPCServer()
	if err != nil {
		log.Fatal(err.Error())
	}
}
