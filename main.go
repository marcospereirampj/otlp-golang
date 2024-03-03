package main

import (
	"context"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
)

const portNum string = ":8080"

func main() {
	log.Println("Starting http server.")

	mux := http.NewServeMux()
	ctx := context.Background()

	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	consoleTraceExporter, err := newTraceExporter()
	if err != nil {
		log.Println("Failed get console exporter (trace).")
	}

	consoleMetricExporter, err := newMetricExporter()
	if err != nil {
		log.Println("Failed get console exporter (metric).")
	}

	tracerProvider := newTraceProvider(consoleTraceExporter)

	//nolint:errcheck
	defer tracerProvider.Shutdown(ctx)
	otel.SetTracerProvider(tracerProvider)

	meterProvider := newMeterProvider(consoleMetricExporter)

	//nolint:errcheck
	defer meterProvider.Shutdown(ctx)
	otel.SetMeterProvider(meterProvider)

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	handleFunc("/info", info)
	newHandler := otelhttp.NewHandler(mux, "/")

	srv := &http.Server{
		Addr:    portNum,
		Handler: newHandler,
	}

	log.Println("Started on port", portNum)
	err = srv.ListenAndServe()
	if err != nil {
		log.Println("Fail start http server.")
	}

}
