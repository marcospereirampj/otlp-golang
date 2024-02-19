package main

import (
	"log"
	"net/http"
)

const portNum string = ":8080"

func main() {
	log.Println("Starting http server.")

	http.HandleFunc("/info", info)

	log.Println("Started on port", portNum)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Fail start http server.")
	}
}
