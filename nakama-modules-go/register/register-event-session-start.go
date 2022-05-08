package register

import (
	"context"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func RegisterEventSessionStart(ctx context.Context, logger runtime.Logger, evt *api.Event) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterEventSessionStart", "event": evt}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
	}
	logger.Info("session start %v %v", ctx, evt)
}
