package after

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/util"
)

func AddFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddFriendsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "AddFriends", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AddGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AddGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "AddGroupUsers", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateAppleRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateCustomRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateDeviceRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateEmailRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateFacebookInstantGameRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGameCenterRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateGoogleRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func AuthenticateSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.AuthenticateSteamRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func BanGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BanGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "BanGroupUsers", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func BlockFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.BlockFriendsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "BlockFriends", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func CreateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Group, in *api.CreateGroupRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DeleteFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteFriendsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DeleteFriends", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DeleteGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteGroupRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DeleteGroup", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DeleteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteLeaderboardRecordRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DeleteLeaderboardRecord", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DeleteNotification(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteNotificationsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DeleteNotification", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DeleteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DeleteStorageObjectsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DeleteStorageObjects", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func DemoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.DemoteGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "DemoteGroupUsers", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func GetAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Account) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "out": out}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func GetUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Users, in *api.GetUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ImportFacebookFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportFacebookFriendsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ImportFacebookFriends", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ImportSteamFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.ImportSteamFriendsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ImportSteamFriends", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func JoinGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinGroupRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "JoinGroup", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func JoinTournament(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.JoinTournamentRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "JoinTournament", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func KickGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.KickGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "KickGroupUsers", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LeaveGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LeaveGroupRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LeaveGroup", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkApple", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkCustom", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkDevice", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkEmail", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkFacebookRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkFacebook", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkFacebookInstantGame", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkGameCenter", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkGoogle", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func LinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.LinkSteamRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "LinkSteam", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListChannelMessages(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ChannelMessageList, in *api.ListChannelMessagesRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListFriends(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.FriendList) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "out": out}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupUserList, in *api.ListGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.GroupList, in *api.ListGroupsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListLeaderboardRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListLeaderboardRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecordList, in *api.ListLeaderboardRecordsAroundOwnerRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListMatches(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.MatchList, in *api.ListMatchesRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListNotifications(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.NotificationList, in *api.ListNotificationsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectList, in *api.ListStorageObjectsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListTournamentRecords(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListTournamentRecordsAroundOwner(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentRecordList, in *api.ListTournamentRecordsAroundOwnerRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListTournaments(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.TournamentList, in *api.ListTournamentsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ListUserGroups(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.UserGroupList, in *api.ListUserGroupsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func PromoteGroupUsers(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.PromoteGroupUsersRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "PromoteGroupUsers", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ReadStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjects, in *api.ReadStorageObjectsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func SessionLogout(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.SessionLogoutRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "SessionLogout", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func SessionRefresh(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.Session, in *api.SessionRefreshRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountApple) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkApple", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkCustom(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountCustom) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkCustom", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkDevice(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountDevice) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkDevice", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkEmail(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountEmail) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkEmail", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkFacebook(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebook) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkFacebook", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkFacebookInstantGame(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountFacebookInstantGame) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkFacebookInstantGame", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkGameCenter(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGameCenter) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkGameCenter", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountGoogle) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkGoogle", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UnlinkSteam(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.AccountSteam) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UnlinkSteam", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UpdateAccount(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateAccountRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UpdateAccount", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func UpdateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.UpdateGroupRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "UpdateGroup", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ValidatePurchaseApple(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseAppleRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ValidatePurchaseGoogle(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseGoogleRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func ValidatePurchaseHuawei(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.ValidatePurchaseResponse, in *api.ValidatePurchaseHuaweiRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func WriteLeaderboardRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteLeaderboardRecordRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func WriteStorageObjects(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.StorageObjectAcks, in *api.WriteStorageObjectsRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
func WriteTournamentRecord(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, out *api.LeaderboardRecord, in *api.WriteTournamentRecordRequest) error {
	if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "ChannelJoin", "in": in}); err != nil {
		logger.Error("Error calling redpanda: %v", err)
		return err
	}
	return nil
}
