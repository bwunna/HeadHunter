package main

import (
	"HeadHunter/internal/delivery"
	"fmt"
)

func main() {
	err := delivery.RunGRPCServer()
	if err != nil {
		fmt.Println(err.Error())
	}

}
