package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeChannelMessageSend(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *rtapi.Envelope) (*rtapi.Envelope, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeChannelMessageSend",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeChannelMessageSend", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
