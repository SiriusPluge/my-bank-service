package main

import (
	mybank "github.com/SiriusPluge/my-bank-service"
	"log"
)

func main() {

	srv := new(mybank.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
