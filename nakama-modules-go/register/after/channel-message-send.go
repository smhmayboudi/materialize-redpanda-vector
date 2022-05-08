package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/rtapi"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func ChannelMessageSend(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out, in *rtapi.Envelope) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelMessageSend", "out": out, "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
