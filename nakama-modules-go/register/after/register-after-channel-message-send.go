package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterAfterChannelMessageSend(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out, in *rtapi.Envelope) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterChannelMessageSend",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelMessageSend", "out": out, "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
