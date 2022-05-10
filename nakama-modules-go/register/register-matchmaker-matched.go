package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterMatchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterMatchmakerMatched",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterMatchmakerMatched", "entries": entries}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return "", err
	}
	return "", nil
}
