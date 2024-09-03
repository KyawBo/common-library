package trace

import (
	"context"

	"github.com/KyawBo/common-library/logger"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func CreateSpan(ctx context.Context, spanName string) (context.Context, trace.Span) {
	return tracer.Start(ctx, spanName)
}

func SetupTracing(ctx context.Context, env string, appName string) *sdktrace.TracerProvider {

	res, err := resource.New(ctx,
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(appName),
		),
	)
	if err != nil {
		logger.Fatalf(ctx, "resource.New: %v", err)
	}
	// Create trace provider without the exporter.
	sampler := sdktrace.ParentBased(
		sdktrace.AlwaysSample(),
		sdktrace.WithRemoteParentSampled(sdktrace.AlwaysSample()),
	)
	var tp *sdktrace.TracerProvider
	if env == "dev" {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithResource(res),
			sdktrace.WithSampler(sampler),
		)
	} else {
		tp = sdktrace.NewTracerProvider(
			sdktrace.WithResource(res),
			sdktrace.WithSampler(sampler),
		)
	}
	// Register the global Tracer provider
	otel.SetTracerProvider(tp)
	tracer = otel.GetTracerProvider().Tracer("example.com/trace")

	return tp
}
