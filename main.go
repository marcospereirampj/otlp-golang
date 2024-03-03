package main

import (
	"log"
	"net/http"
)

const portNum string = ":8080"

func main() {
	log.Println("Starting http server.")
	log.Println("Started on port", portNum)

	mux := http.NewServeMux()

	mux.HandleFunc("/info", info)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Println("Fail start http server.")
	}

}
