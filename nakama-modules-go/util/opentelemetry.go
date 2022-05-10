// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/otel-collector/main.go
package util

import (
	"context"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

const (
	serviceName    = "nakama-modules-go"
	serviceVersion = "v0.1.0"
	url            = "http://jaeger:14268/api/traces"
)

func InitProvider(ctx context.Context, logger runtime.Logger) func() {
	// b3 := b3.New(b3.WithInjectEncoding(b3.B3SingleHeader))
	// otel.SetTextMapPropagator(b3)

	rna := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(serviceVersion),
	)
	rm, err := resource.Merge(
		resource.Default(),
		rna,
	)
	if err != nil {
		logger.Error("Failed to create resource: %v", err)
	}
	rn, err := resource.New(
		ctx,
		// resource.WithAttributes(),
		resource.WithContainer(),
		// resource.WithDetectors(thirdparty.Detector{}),
		resource.WithFromEnv(),
		resource.WithOS(),
		resource.WithProcess(),
	)
	r, err := resource.Merge(
		rm,
		rn,
	)
	if err != nil {
		logger.Error("Failed to create resource: %v", err)
	}
	ej, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		logger.Error("Failed to create jaeger exporter: %v", err)
	}
	es, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		logger.Error("Failed to create stdouttrace exporter: %v", err)
	}
	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(ej),
		sdkTrace.WithBatcher(es),
		sdkTrace.WithResource(r),
	)
	otel.SetTracerProvider(tp)

	return func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer func(ctx context.Context) {
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := tp.Shutdown(ctx); err != nil {
				logger.Error("Failed to shutdown tracer provider: %v", err)
			}
		}(ctx)
	}
}
