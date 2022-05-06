package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func RegisterMatchmakerMatched(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, entries []runtime.MatchmakerEntry) (string, error) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "registerEvent", "entries": entries}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return "", err
	}
	return "", nil
}
