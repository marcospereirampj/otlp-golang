package main

import (
	"log"
	"net/http"
)

const portNum string = ":8080"

func main() {
	log.Println("Starting http server.")

	mux := http.NewServeMux()

	mux.HandleFunc("/info", info)

	srv := &http.Server{
		Addr:    portNum,
		Handler: mux,
	}

	log.Println("Started on port", portNum)
	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Fail start http server.")
	}

}
