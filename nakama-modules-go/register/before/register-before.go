package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeAddFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) (*api.AddFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAddFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAddFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAddGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) (*api.AddGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAddGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAddGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateAppleRequest) (*api.AuthenticateAppleRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateApple", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateCustomRequest) (*api.AuthenticateCustomRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateCustom", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateDeviceRequest) (*api.AuthenticateDeviceRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateDevice", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateEmailRequest) (*api.AuthenticateEmailRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateEmail", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateFacebookRequest) (*api.AuthenticateFacebookRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateFacebook", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateFacebookInstantGameRequest) (*api.AuthenticateFacebookInstantGameRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateFacebookInstantGame", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateGameCenterRequest) (*api.AuthenticateGameCenterRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateGameCenter", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateGoogleRequest) (*api.AuthenticateGoogleRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateGoogle", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeAuthenticateSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AuthenticateSteamRequest) (*api.AuthenticateSteamRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeAuthenticateSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeAuthenticateSteam", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeBanGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BanGroupUsersRequest) (*api.BanGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeBanGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeBanGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeBlockFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) (*api.BlockFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeBlockFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeBlockFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeCreateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeCreateGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeCreateGroup", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDeleteFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) (*api.DeleteFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDeleteFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDeleteFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDeleteGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) (*api.DeleteGroupRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDeleteGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDeleteGroup", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDeleteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteLeaderboardRecordRequest) (*api.DeleteLeaderboardRecordRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDeleteLeaderboardRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDeleteLeaderboardRecord", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDeleteNotification(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteNotificationsRequest) (*api.DeleteNotificationsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDeleteNotification",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDeleteNotification", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDeleteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteStorageObjectsRequest) (*api.DeleteStorageObjectsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDeleteStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDeleteStorageObjects", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeDemoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DemoteGroupUsersRequest) (*api.DemoteGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeDemoteGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeDemoteGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeGetAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeGetAccount",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeGetAccount"}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return err
	// }
	return nil
}
func RegisterBeforeGetUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.GetUsersRequest) (*api.GetUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeGetUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeGetUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeImportFacebookFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportFacebookFriendsRequest) (*api.ImportFacebookFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeImportFacebookFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeImportFacebookFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeImportSteamFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportSteamFriendsRequest) (*api.ImportSteamFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeImportSteamFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeImportSteamFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeJoinGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) (*api.JoinGroupRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeJoinGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeJoinGroup", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeJoinTournament(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinTournamentRequest) (*api.JoinTournamentRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeJoinTournament",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeJoinTournament", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeKickGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) (*api.KickGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeKickGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeKickGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLeaveGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) (*api.LeaveGroupRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLeaveGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLeaveGroup", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) (*api.AccountApple, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkApple", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkCustom", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkDevice", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkEmail", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) (*api.LinkFacebookRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkFacebook", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkFacebookInstantGame", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkGameCenter", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkGoogle", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeLinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkSteamRequest) (*api.LinkSteamRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeLinkSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeLinkSteam", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListChannelMessages(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListChannelMessagesRequest) (*api.ListChannelMessagesRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListChannelMessages",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListChannelMessages", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListFriendsRequest) (*api.ListFriendsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListFriends", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListGroupUsersRequest) (*api.ListGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListGroupsRequest) (*api.ListGroupsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListGroups",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListGroups", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListLeaderboardRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListLeaderboardRecordsRequest) (*api.ListLeaderboardRecordsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListLeaderboardRecords",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListLeaderboardRecords", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListLeaderboardRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.ListLeaderboardRecordsAroundOwnerRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListLeaderboardRecordsAroundOwner",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListLeaderboardRecordsAroundOwner", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListMatches(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListMatchesRequest) (*api.ListMatchesRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListMatches",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListMatches", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListNotifications(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListNotificationsRequest) (*api.ListNotificationsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListNotifications",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListNotifications", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListStorageObjectsRequest) (*api.ListStorageObjectsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListStorageObjects", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListTournamentRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentRecordsRequest) (*api.ListTournamentRecordsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListTournamentRecords",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListTournamentRecords", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListTournamentRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentRecordsAroundOwnerRequest) (*api.ListTournamentRecordsAroundOwnerRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListTournamentRecordsAroundOwner",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListTournamentRecordsAroundOwner", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListTournaments(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListTournamentsRequest) (*api.ListTournamentsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListTournaments",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListTournaments", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeListUserGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ListUserGroupsRequest) (*api.ListUserGroupsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeListUserGroups",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeListUserGroups", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforePromoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) (*api.PromoteGroupUsersRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforePromoteGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforePromoteGroupUsers", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeReadStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ReadStorageObjectsRequest) (*api.ReadStorageObjectsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeReadStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeReadStorageObjects", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeSessionLogout(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.SessionLogoutRequest) (*api.SessionLogoutRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeSessionLogout",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeSessionLogout", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeSessionRefresh(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.SessionRefreshRequest) (*api.SessionRefreshRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeSessionRefresh",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeSessionRefresh", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) (*api.AccountApple, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkApple", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) (*api.AccountCustom, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkCustom", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) (*api.AccountDevice, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkDevice", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) (*api.AccountEmail, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkEmail", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebook) (*api.AccountFacebook, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkFacebook", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) (*api.AccountFacebookInstantGame, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkFacebookInstantGame", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) (*api.AccountGameCenter, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkGameCenter", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) (*api.AccountGoogle, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkGoogle", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUnlinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) (*api.AccountSteam, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUnlinkSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUnlinkSteam", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUpdateAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateAccountRequest) (*api.UpdateAccountRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUpdateAccount",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUpdateAccount", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeUpdateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) (*api.UpdateGroupRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeUpdateGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeUpdateGroup", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeValidatePurchaseApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ValidatePurchaseAppleRequest) (*api.ValidatePurchaseAppleRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeValidatePurchaseApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeValidatePurchaseApple", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeValidatePurchaseGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ValidatePurchaseGoogleRequest) (*api.ValidatePurchaseGoogleRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeValidatePurchaseGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeValidatePurchaseGoogle", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeValidatePurchaseHuawei(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ValidatePurchaseHuaweiRequest) (*api.ValidatePurchaseHuaweiRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeValidatePurchaseHuawei",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeValidatePurchaseHuawei", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeWriteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteLeaderboardRecordRequest) (*api.WriteLeaderboardRecordRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeWriteLeaderboardRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeWriteLeaderboardRecord", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeWriteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteStorageObjectsRequest) (*api.WriteStorageObjectsRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeWriteStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeWriteStorageObjects", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
func RegisterBeforeWriteTournamentRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.WriteTournamentRecordRequest) (*api.WriteTournamentRecordRequest, error) {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterBeforeWriteTournamentRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeWriteTournamentRecord", "in": in}); err != nil {
	// 	logger.Error("Error calling redpanda: %v", err)
	// 	return in, err
	// }
	return in, nil
}
