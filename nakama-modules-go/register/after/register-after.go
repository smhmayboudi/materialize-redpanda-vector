package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

func RegisterAfterAddFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAddFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterAddFriends", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAddGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAddGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterAddGroupUsers", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateAppleRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateCustomRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateDeviceRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGameCenterRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGoogleRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterAuthenticateSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateSteamRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterAuthenticateSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterBanGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BanGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterBanGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterBanGroupUsers", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterBlockFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterBlockFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterBlockFriends", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterCreateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Group, in *api.CreateGroupRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterCreateGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDeleteFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDeleteFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDeleteFriends", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDeleteGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDeleteGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDeleteGroup", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDeleteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteLeaderboardRecordRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDeleteLeaderboardRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDeleteLeaderboardRecord", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDeleteNotification(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteNotificationsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDeleteNotification",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDeleteNotification", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDeleteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteStorageObjectsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDeleteStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDeleteStorageObjects", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterDemoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DemoteGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterDemoteGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterDemoteGroupUsers", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterGetAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Account) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterGetAccount",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "out": out}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterGetUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Users, in *api.GetUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterGetUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterImportFacebookFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportFacebookFriendsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterImportFacebookFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterImportFacebookFriends", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterImportSteamFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportSteamFriendsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterImportSteamFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterImportSteamFriends", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterJoinGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterJoinGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterJoinGroup", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterJoinTournament(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinTournamentRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterJoinTournament",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterJoinTournament", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterKickGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterKickGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterKickGroupUsers", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLeaveGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLeaveGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLeaveGroup", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkApple", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkCustom", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkDevice", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkEmail", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkFacebook", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkFacebookInstantGame", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkGameCenter", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkGoogle", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterLinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkSteamRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterLinkSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterLinkSteam", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListChannelMessages(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListChannelMessages",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.FriendList) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListFriends",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "out": out}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupUserList, in *api.ListGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupList, in *api.ListGroupsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListGroups",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListLeaderboardRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListLeaderboardRecords",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListLeaderboardRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListLeaderboardRecordsAroundOwner",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListMatches(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.MatchList, in *api.ListMatchesRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListMatches",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListNotifications(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.NotificationList, in *api.ListNotificationsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListNotifications",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListTournamentRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListTournamentRecords",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListTournamentRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListTournamentRecordsAroundOwner",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListTournaments(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentList, in *api.ListTournamentsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListTournaments",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterListUserGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.UserGroupList, in *api.ListUserGroupsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterListUserGroups",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterPromoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterPromoteGroupUsers",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterPromoteGroupUsers", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterReadStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterReadStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterSessionLogout(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.SessionLogoutRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterSessionLogout",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterSessionLogout", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterSessionRefresh(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.SessionRefreshRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterSessionRefresh",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkApple", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkCustom",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkCustom", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkDevice",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkDevice", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkEmail",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkEmail", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebook) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkFacebook",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkFacebook", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkFacebookInstantGame",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkFacebookInstantGame", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkGameCenter",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkGameCenter", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkGoogle", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUnlinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUnlinkSteam",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUnlinkSteam", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUpdateAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateAccountRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUpdateAccount",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUpdateAccount", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterUpdateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterUpdateGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterUpdateGroup", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterValidatePurchaseApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseAppleRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterValidatePurchaseApple",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterValidatePurchaseGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseGoogleRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterValidatePurchaseGoogle",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterValidatePurchaseHuawei(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseHuaweiRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterValidatePurchaseHuawei",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterWriteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterWriteLeaderboardRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterWriteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterWriteStorageObjects",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
func RegisterAfterWriteTournamentRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error {
	ctx, span := otel.Tracer(u.InstrumentationName).Start(
		ctx,
		"RegisterAfterWriteTournamentRecord",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterAfterChannelJoin", "in": in}); err != nil {
		textMapCarrier := u.NewTextMapCarrier(ctx)
		logger.WithFields(textMapCarrier.Fields()).WithField("error", err).Error("Error calling redpanda")
		span.RecordError(err)
		span.SetStatus(codes.Error, "Error calling redpanda")
		return err
	}
	return nil
}
