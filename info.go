package main

import (
	"encoding/json"
	"net/http"
)

type InfoResponse struct {
	Version     string `json:"version"`
	ServiceName string `json:"service-name"`
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := InfoResponse{Version: "0.1.0", ServiceName: "otlp-sample"}
	//nolint:errcheck
	json.NewEncoder(w).Encode(response)
}
