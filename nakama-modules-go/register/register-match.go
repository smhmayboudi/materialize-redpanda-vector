// https://github.com/heroiclabs/nakama/blob/master/sample_go_module/sample.go
package register

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

type Match struct{}

type MatchState struct {
	debug bool
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchInit", "params": params}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return params, 0, ""
	}
	var debug bool
	if d, ok := params["debug"]; ok {
		if dv, ok := d.(bool); ok {
			debug = dv
		}
	}
	state := &MatchState{
		debug: debug,
	}
	if state.debug {
		logger.Info("match init, starting with debug: %v", state.debug)
	}
	tickRate := 1
	label := "skill=100-150"
	return state, tickRate, label
}
func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchJoinAttempt", "tick": tick, "state": state, "presence": presence, "metadata": metadata}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state, false, ""
	}
	if state.(*MatchState).debug {
		logger.Info("match join attempt username %v user_id %v session_id %v node %v with metadata %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId(), metadata)
	}
	return state, true, ""
}
func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchJoin", "tick": tick, "state": state, "presences": presences}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state
	}
	if state.(*MatchState).debug {
		for _, presence := range presences {
			logger.Info("match join username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
		}
	}
	return state
}
func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchLeave", "tick": tick, "state": state, "presences": presences}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state
	}
	if state.(*MatchState).debug {
		for _, presence := range presences {
			logger.Info("match leave username %v user_id %v session_id %v node %v", presence.GetUsername(), presence.GetUserId(), presence.GetSessionId(), presence.GetNodeId())
		}
	}
	return state
}
func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchLeave", "tick": tick, "state": state, "messages": messages}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state
	}
	if state.(*MatchState).debug {
		logger.Info("match loop match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Info("match loop match_id %v message count %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), len(messages))
	}
	if tick >= 10 {
		return nil
	}
	return state
}
func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchTerminate", "tick": tick, "state": state, "graceSeconds": graceSeconds}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state
	}
	if state.(*MatchState).debug {
		logger.Info("match terminate match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Info("match terminate match_id %v grace seconds %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), graceSeconds)
	}
	return state
}
func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "matchSignal", "tick": tick, "state": state, "data": data}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return state, data
	}
	if state.(*MatchState).debug {
		logger.Info("match signal match_id %v tick %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), tick)
		logger.Info("match signal match_id %v data %v", ctx.Value(runtime.RUNTIME_CTX_MATCH_ID), data)
	}
	return state, data
}

func RegisterMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
	return &Match{}, nil
}
