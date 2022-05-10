package register

import (
	"context"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterEventSessionEnd(ctx context.Context, logger runtime.Logger, evt *api.Event) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterEventSessionEnd",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterEventSessionEnd", "event": evt}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
	}
	logger.Info("session end %v %v", ctx, evt)
}
