package main

import (
	"log"
	"net/http"
	"restapi/src/routes"
)

func main() {
	routes.Setup()

	err := http.ListenAndServe(":10228", nil)
	
	if err != nil { 
		log.Fatal(err.Error())
	}	
}