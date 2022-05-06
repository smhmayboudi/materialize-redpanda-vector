package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	r "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/register"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterLeaderboardReset(r.RegisterLeaderboardReset); err != nil {
		return err
	}
	if err := initializer.RegisterMatch("", r.RegisterMatch); err != nil {
		return err
	}
	if err := initializer.RegisterMatchmakerMatched(r.RegisterMatchmakerMatched); err != nil {
		return err
	}
	if err := initializer.RegisterTournamentEnd(r.RegisterTournamentEnd); err != nil {
		return err
	}
	if err := initializer.RegisterTournamentReset(r.RegisterTournamentReset); err != nil {
		return err
	}
	if err := initializer.RegisterEvent(r.RegisterEvent); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionStart(r.RegisterEventSessionStart); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionEnd(r.RegisterEventSessionEnd); err != nil {
		return err
	}
	return nil
}
