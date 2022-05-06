package util

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/heroiclabs/nakama-common/runtime"
)

type NakamaContext struct {
	ClientIp       string            `json:"client_ip,omitempty"`
	ClientSort     string            `json:"client_port,omitempty"`
	Env            map[string]string `json:"env,omitempty"`
	ExecutionMode  string            `json:"execution_mode,omitempty"`
	Headers        map[string]string `json:"headers,omitempty"`
	Lang           string            `json:"lang,omitempty"`
	MatchId        string            `json:"match_id,omitempty"`
	MatchLabel     string            `json:"match_label,omitempty"`
	MatchNode      string            `json:"match_node,omitempty"`
	MatchTickRate  int               `json:"match_tick_rate,omitempty"`
	Node           string            `json:"node,omitempty"`
	QueryParams    map[string]string `json:"query_params,omitempty"`
	SessionId      string            `json:"session_id,omitempty"`
	UserId         string            `json:"user_id,omitempty"`
	UserSessionExp int               `json:"user_session_exp,omitempty"`
	Username       string            `json:"username,omitempty"`
	Vars           map[string]string `json:"vars,omitempty"`
}

type DataKey struct {
	Node string `json:"node"`
}

type DataValue struct {
	Context NakamaContext          `json:"context,omitempty"`
	Payload map[string]interface{} `json:"payload,omitempty"`
}

type Record struct {
	Key       DataKey   `json:"key,omitempty"`
	Partition int       `json:"partition,omitempty"`
	Value     DataValue `json:"value,omitempty"`
}

type Records struct {
	Records []Record `json:"records"`
}

func Redpanda(ctx context.Context, logger runtime.Logger, payload map[string]interface{}) error {
	env, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string)
	if !ok {
		env = map[string]string{}
	}
	executionMode, ok := ctx.Value(runtime.RUNTIME_CTX_MODE).(string)
	if !ok {
		executionMode = ""
	}
	node, ok := ctx.Value(runtime.RUNTIME_CTX_NODE).(string)
	if !ok {
		node = ""
	}
	headers, ok := ctx.Value(runtime.RUNTIME_CTX_HEADERS).(map[string]string)
	if !ok {
		headers = map[string]string{}
	}
	dataKey := &DataKey{
		Node: node,
	}
	queryParams, ok := ctx.Value(runtime.RUNTIME_CTX_QUERY_PARAMS).(map[string]string)
	if !ok {
		queryParams = map[string]string{}
	}
	userId, ok := ctx.Value(runtime.RUNTIME_CTX_USER_ID).(string)
	if !ok {
		userId = ""
	}
	username, ok := ctx.Value(runtime.RUNTIME_CTX_USERNAME).(string)
	if !ok {
		userId = ""
	}
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)
	if !ok {
		vars = map[string]string{}
	}
	userSessionExp, ok := ctx.Value(runtime.RUNTIME_CTX_USER_SESSION_EXP).(int)
	if !ok {
		userSessionExp = 0
	}
	sessionId, ok := ctx.Value(runtime.RUNTIME_CTX_SESSION_ID).(string)
	if !ok {
		sessionId = ""
	}
	lang, ok := ctx.Value(runtime.RUNTIME_CTX_LANG).(string)
	if !ok {
		lang = ""
	}
	clientIp, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_IP).(string)
	if !ok {
		clientIp = ""
	}
	clientSort, ok := ctx.Value(runtime.RUNTIME_CTX_CLIENT_PORT).(string)
	if !ok {
		clientSort = ""
	}
	matchId, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_ID).(string)
	if !ok {
		matchId = ""
	}
	matchNode, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_NODE).(string)
	if !ok {
		matchNode = ""
	}
	matchLabel, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_LABEL).(string)
	if !ok {
		matchLabel = ""
	}
	matchTickRate, ok := ctx.Value(runtime.RUNTIME_CTX_MATCH_TICK_RATE).(int)
	if !ok {
		matchTickRate = 0
	}
	nakamaContext := &NakamaContext{
		ClientIp:       clientIp,
		ClientSort:     clientSort,
		Env:            env,
		ExecutionMode:  executionMode,
		Headers:        headers,
		Lang:           lang,
		MatchId:        matchId,
		MatchLabel:     matchLabel,
		MatchNode:      matchNode,
		MatchTickRate:  matchTickRate,
		Node:           node,
		QueryParams:    queryParams,
		SessionId:      sessionId,
		UserId:         userId,
		UserSessionExp: userSessionExp,
		Username:       username,
		Vars:           vars,
	}
	dataValue := &DataValue{
		Context: *nakamaContext,
		Payload: payload,
	}
	record := &Record{
		Key:       *dataKey,
		Partition: 0,
		Value:     *dataValue,
	}
	items := make([]Record, 1)
	items[0] = *record
	records := &Records{
		Records: items,
	}
	body, err := json.Marshal(records)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	res, err := http.Post("http://redpanda:8082/topics/nakama", "application/vnd.kafka.json.v2+json", bytes.NewReader(body))
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	defer res.Body.Close()
	return nil
}
