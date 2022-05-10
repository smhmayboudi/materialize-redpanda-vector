package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterTournamentEnd(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, tournament *api.Tournament, end, reset int64) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterTournamentEnd",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterTournamentEnd", "tournament": tournament, "end": end, "reset": reset}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
