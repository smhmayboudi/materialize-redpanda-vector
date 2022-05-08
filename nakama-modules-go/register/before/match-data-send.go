package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func MatchDataSend(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *rtapi.Envelope) (*rtapi.Envelope, error) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "MatchDataSend", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return in, err
	}
	return in, nil
}
