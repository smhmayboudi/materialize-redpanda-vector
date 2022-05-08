package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func RegisterLeaderboardReset(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, leaderboard *api.Leaderboard, reset int64) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterLeaderboardReset", "leaderboard": leaderboard, "reset": reset}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
