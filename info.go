package main

import (
	"encoding/json"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
	"net/http"
)

type InfoResponse struct {
	Version     string `json:"version"`
	ServiceName string `json:"service-name"`
}

var (
	tracer      = otel.Tracer("info-service")
	meter       = otel.Meter("info-service")
	viewCounter metric.Int64Counter
)

func init() {
	var err error
	viewCounter, err = meter.Int64Counter("user.views",
		metric.WithDescription("The number of views"),
		metric.WithUnit("{views}"))
	if err != nil {
		panic(err)
	}
}

func info(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "info")
	defer span.End()

	viewCounter.Add(ctx, 1)

	w.Header().Set("Content-Type", "application/json")
	response := InfoResponse{Version: "0.1.0", ServiceName: "otlp-sample"}
	//nolint:errcheck
	json.NewEncoder(w).Encode(response)
}
