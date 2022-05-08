// https://github.com/open-telemetry/opentelemetry-go/blob/main/example/otel-collector/main.go
package util

import (
	"context"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func initProvider(ctx context.Context, logger runtime.Logger) func() {
	// ctx := context.Background()
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			attribute.String("environment", "development"),
			semconv.ServiceNameKey.String("nakama-modules-go"),
			semconv.ServiceVersionKey.String("v0.1.0"),
		),
	)
	if err != nil {
		logger.Error("Failed to create resource: %v", err)
	}
	traceExporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint("http://localhost:14268/api/traces")))
	if err != nil {
		logger.Error("Failed to create trace exporter: %v", err)
	}
	bsp := sdkTrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdkTrace.NewTracerProvider(
		sdkTrace.WithSampler(sdkTrace.AlwaysSample()),
		sdkTrace.WithResource(r),
		sdkTrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		defer func(ctx context.Context) {
			ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := tracerProvider.Shutdown(ctx); err != nil {
				logger.Error("Failed to shutdown tracer provider: %v", err)
			}
		}(ctx)
		// if err != nil {
		// 	logger.Error("Failed to shutdown tracer provider: %v", err)
		// }
	}
}
