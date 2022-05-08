package main

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/runtime"
	r "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/register"
	ra "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/register/after"
	rb "github.com/smhmayboudi/materialize-redpanda-vector/nakama-modules-go/register/before"
)

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	if err := initializer.RegisterAfterAddFriends(ra.AddFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAddGroupUsers(ra.AddGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateApple(ra.AuthenticateApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateCustom(ra.AuthenticateCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateDevice(ra.AuthenticateDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateEmail(ra.AuthenticateEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateFacebook(ra.AuthenticateFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateFacebookInstantGame(ra.AuthenticateFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateGameCenter(ra.AuthenticateGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateGoogle(ra.AuthenticateGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterAuthenticateSteam(ra.AuthenticateSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterBanGroupUsers(ra.BanGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterBlockFriends(ra.BlockFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterCreateGroup(ra.CreateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteFriends(ra.DeleteFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteGroup(ra.DeleteGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteLeaderboardRecord(ra.DeleteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteNotification(ra.DeleteNotification); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDeleteStorageObjects(ra.DeleteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterDemoteGroupUsers(ra.DemoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterGetAccount(ra.GetAccount); err != nil {
		return err
	}
	if err := initializer.RegisterAfterGetUsers(ra.GetUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterImportFacebookFriends(ra.ImportFacebookFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterImportSteamFriends(ra.ImportSteamFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterJoinGroup(ra.JoinGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterJoinTournament(ra.JoinTournament); err != nil {
		return err
	}
	if err := initializer.RegisterAfterKickGroupUsers(ra.KickGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLeaveGroup(ra.LeaveGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkApple(ra.LinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkCustom(ra.LinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkDevice(ra.LinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkEmail(ra.LinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkFacebook(ra.LinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkFacebookInstantGame(ra.LinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkGameCenter(ra.LinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkGoogle(ra.LinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterLinkSteam(ra.LinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListChannelMessages(ra.ListChannelMessages); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListFriends(ra.ListFriends); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListGroupUsers(ra.ListGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListGroups(ra.ListGroups); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListLeaderboardRecords(ra.ListLeaderboardRecords); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListLeaderboardRecordsAroundOwner(ra.ListLeaderboardRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListMatches(ra.ListMatches); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListNotifications(ra.ListNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListStorageObjects(ra.ListStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournamentRecords(ra.ListTournamentRecords); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournamentRecordsAroundOwner(ra.ListTournamentRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListTournaments(ra.ListTournaments); err != nil {
		return err
	}
	if err := initializer.RegisterAfterListUserGroups(ra.ListUserGroups); err != nil {
		return err
	}
	if err := initializer.RegisterAfterPromoteGroupUsers(ra.PromoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterAfterReadStorageObjects(ra.ReadStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterSessionLogout(ra.SessionLogout); err != nil {
		return err
	}
	if err := initializer.RegisterAfterSessionRefresh(ra.SessionRefresh); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkApple(ra.UnlinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkCustom(ra.UnlinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkDevice(ra.UnlinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkEmail(ra.UnlinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkFacebook(ra.UnlinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkFacebookInstantGame(ra.UnlinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkGameCenter(ra.UnlinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkGoogle(ra.UnlinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUnlinkSteam(ra.UnlinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUpdateAccount(ra.UpdateAccount); err != nil {
		return err
	}
	if err := initializer.RegisterAfterUpdateGroup(ra.UpdateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseApple(ra.ValidatePurchaseApple); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseGoogle(ra.ValidatePurchaseGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterAfterValidatePurchaseHuawei(ra.ValidatePurchaseHuawei); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteLeaderboardRecord(ra.WriteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteStorageObjects(ra.WriteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterAfterWriteTournamentRecord(ra.WriteTournamentRecord); err != nil {
		return err
	}

	if err := initializer.RegisterBeforeAddFriends(rb.AddFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAddGroupUsers(rb.AddGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateApple(rb.AuthenticateApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateCustom(rb.AuthenticateCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateDevice(rb.AuthenticateDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateEmail(rb.AuthenticateEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateFacebook(rb.AuthenticateFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateFacebookInstantGame(rb.AuthenticateFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateGameCenter(rb.AuthenticateGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateGoogle(rb.AuthenticateGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeAuthenticateSteam(rb.AuthenticateSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeBanGroupUsers(rb.BanGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeBlockFriends(rb.BlockFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeCreateGroup(rb.CreateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteFriends(rb.DeleteFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteGroup(rb.DeleteGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteLeaderboardRecord(rb.DeleteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteNotification(rb.DeleteNotification); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDeleteStorageObjects(rb.DeleteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeDemoteGroupUsers(rb.DemoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeGetAccount(rb.GetAccount); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeGetUsers(rb.GetUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeImportFacebookFriends(rb.ImportFacebookFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeImportSteamFriends(rb.ImportSteamFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeJoinGroup(rb.JoinGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeJoinTournament(rb.JoinTournament); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeKickGroupUsers(rb.KickGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLeaveGroup(rb.LeaveGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkApple(rb.LinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkCustom(rb.LinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkDevice(rb.LinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkEmail(rb.LinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkFacebook(rb.LinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkFacebookInstantGame(rb.LinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkGameCenter(rb.LinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkGoogle(rb.LinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeLinkSteam(rb.LinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListChannelMessages(rb.ListChannelMessages); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListFriends(rb.ListFriends); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListGroupUsers(rb.ListGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListGroups(rb.ListGroups); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListLeaderboardRecords(rb.ListLeaderboardRecords); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListLeaderboardRecordsAroundOwner(rb.ListLeaderboardRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListMatches(rb.ListMatches); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListNotifications(rb.ListNotifications); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListStorageObjects(rb.ListStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournamentRecords(rb.ListTournamentRecords); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournamentRecordsAroundOwner(rb.ListTournamentRecordsAroundOwner); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListTournaments(rb.ListTournaments); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeListUserGroups(rb.ListUserGroups); err != nil {
		return err
	}
	if err := initializer.RegisterBeforePromoteGroupUsers(rb.PromoteGroupUsers); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeReadStorageObjects(rb.ReadStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeSessionLogout(rb.SessionLogout); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeSessionRefresh(rb.SessionRefresh); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkApple(rb.UnlinkApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkCustom(rb.UnlinkCustom); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkDevice(rb.UnlinkDevice); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkEmail(rb.UnlinkEmail); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkFacebook(rb.UnlinkFacebook); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkFacebookInstantGame(rb.UnlinkFacebookInstantGame); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkGameCenter(rb.UnlinkGameCenter); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkGoogle(rb.UnlinkGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUnlinkSteam(rb.UnlinkSteam); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUpdateAccount(rb.UpdateAccount); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeUpdateGroup(rb.UpdateGroup); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseApple(rb.ValidatePurchaseApple); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseGoogle(rb.ValidatePurchaseGoogle); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeValidatePurchaseHuawei(rb.ValidatePurchaseHuawei); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteLeaderboardRecord(rb.WriteLeaderboardRecord); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteStorageObjects(rb.WriteStorageObjects); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeWriteTournamentRecord(rb.WriteTournamentRecord); err != nil {
		return err
	}

	if err := initializer.RegisterAfterRt("ChannelJoin", ra.ChannelJoin); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelLeave", ra.ChannelLeave); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageRemove", ra.ChannelMessageRemove); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageSend", ra.ChannelMessageSend); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("ChannelMessageUpdate", ra.ChannelMessageUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchCreate", ra.MatchCreate); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchDataSend", ra.MatchDataSend); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchJoin", ra.MatchJoin); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchLeave", ra.MatchLeave); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchmakerAdd", ra.MatchmakerAdd); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("MatchmakerRemove", ra.MatchmakerRemove); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("Ping", ra.Ping); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("Pong", ra.Pong); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusFollow", ra.StatusFollow); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusUnfollow", ra.StatusUnfollow); err != nil {
		return err
	}
	if err := initializer.RegisterAfterRt("StatusUpdate", ra.StatusUpdate); err != nil {
		return err
	}

	if err := initializer.RegisterBeforeRt("ChannelJoin", rb.ChannelJoin); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelLeave", rb.ChannelLeave); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageRemove", rb.ChannelMessageRemove); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageSend", rb.ChannelMessageSend); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("ChannelMessageUpdate", rb.ChannelMessageUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchCreate", rb.MatchCreate); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchDataSend", rb.MatchDataSend); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchJoin", rb.MatchJoin); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchLeave", rb.MatchLeave); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchmakerAdd", rb.MatchmakerAdd); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("MatchmakerRemove", rb.MatchmakerRemove); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("Ping", rb.Ping); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("Pong", rb.Pong); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusFollow", rb.StatusFollow); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusUnfollow", rb.StatusUnfollow); err != nil {
		return err
	}
	if err := initializer.RegisterBeforeRt("StatusUpdate", rb.StatusUpdate); err != nil {
		return err
	}
	if err := initializer.RegisterEvent(r.RegisterEvent); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionEnd(r.RegisterEventSessionEnd); err != nil {
		return err
	}
	if err := initializer.RegisterEventSessionStart(r.RegisterEventSessionStart); err != nil {
		return err
	}
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
	return nil
}
