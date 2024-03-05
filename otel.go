package main

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"time"

	mexporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric"
	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
)

//nolint:all
func newTraceExporter() (trace.SpanExporter, error) {
	return stdouttrace.New(stdouttrace.WithPrettyPrint())
}

//nolint:all
func newMetricExporter() (metric.Exporter, error) {
	return stdoutmetric.New()
}

func newMetricGoogleExporter(projectID string) (metric.Exporter, error) {
	return mexporter.New(mexporter.WithProjectID(projectID))
}

func newTraceGoogleExporter(projectID string) (trace.SpanExporter, error) {
	return texporter.New(texporter.WithProjectID(projectID))
}

func newTraceProvider(traceExporter trace.SpanExporter) *trace.TracerProvider {
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(traceExporter,
			trace.WithBatchTimeout(time.Second)),
	)
	return traceProvider
}

func newMeterProvider(meterExporter metric.Exporter) *metric.MeterProvider {
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(meterExporter,
			metric.WithInterval(10*time.Second))),
	)
	return meterProvider
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	)
}
