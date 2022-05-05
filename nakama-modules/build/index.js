"use strict";
var afterAddFriends = function (ctx, logger, nk, data, request) { };
var afterAddGroupUsers = function (ctx, logger, nk, data, request) { };
var afterAuthenticateApple = function (ctx, logger, nk, data, request) { };
var afterAuthenticateCustom = function (ctx, logger, nk, data, request) { };
var afterAuthenticateDevice = function (ctx, logger, nk, data, request) { };
var afterAuthenticateEmail = function (ctx, logger, nk, data, request) { };
var afterAuthenticateFacebook = function (ctx, logger, nk, data, request) { };
var afterAuthenticateFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var afterAuthenticateGameCenter = function (ctx, logger, nk, data, request) { };
var afterAuthenticateGoogle = function (ctx, logger, nk, data, request) { };
var afterAuthenticateSteam = function (ctx, logger, nk, data, request) { };
var afterBanGroupUsers = function (ctx, logger, nk, data, request) { };
var afterBlockFriends = function (ctx, logger, nk, data, request) { };
var afterChannelJoin = function (ctx, logger, nk, output, input) {
    if (typeof input.channelJoin !==
        "undefined" &&
        output !== null &&
        typeof output.channelJoin !== "undefined") {
    }
};
var afterChannelLeave = function (ctx, logger, nk, output, input) {
    if (typeof input.channelLeave !==
        "undefined" &&
        output !== null &&
        typeof output.channelLeave !==
            "undefined") {
    }
};
var afterChannelMessageRemove = function (ctx, logger, nk, output, input) {
    if (typeof input
        .channelMessageRemove !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageRemove !== "undefined") {
    }
};
var afterChannelMessageSend = function (ctx, logger, nk, output, input) {
    if (typeof input
        .channelMessageSend !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageSend !== "undefined") {
    }
};
var afterChannelMessageUpdate = function (ctx, logger, nk, output, input) {
    if (typeof input
        .channelMessageUpdate !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageUpdate !== "undefined") {
    }
};
var afterCreateGroup = function (ctx, logger, nk, data, request) { };
var afterDeleteFriends = function (ctx, logger, nk, data, request) { };
var afterDeleteGroup = function (ctx, logger, nk, data, request) { };
var afterDeleteLeaderboardRecord = function (ctx, logger, nk, data, request) { };
var afterDeleteNotifications = function (ctx, logger, nk, data, request) { };
var afterDeleteStorageObjects = function (ctx, logger, nk, data, request) { };
var afterDemoteGroupUsers = function (ctx, logger, nk, data, request) { };
var afterEvent = function (ctx, logger, nk, data, request) { };
var afterGetAccount = function (ctx, logger, nk, data, request) { };
var afterGetUsers = function (ctx, logger, nk, data, request) { };
var afterImportFacebookFriends = function (ctx, logger, nk, data, request) { };
var afterImportSteamFriends = function (ctx, logger, nk, data, request) { };
var afterJoinGroup = function (ctx, logger, nk, data, request) { };
var afterJoinTournament = function (ctx, logger, nk, data, request) { };
var afterKickGroupUsers = function (ctx, logger, nk, data, request) { };
var afterLeaveGroup = function (ctx, logger, nk, data, request) { };
var afterLinkApple = function (ctx, logger, nk, data, request) { };
var afterLinkCustom = function (ctx, logger, nk, data, request) { };
var afterLinkDevice = function (ctx, logger, nk, data, request) { };
var afterLinkEmail = function (ctx, logger, nk, data, request) { };
var afterLinkFacebook = function (ctx, logger, nk, data, request) { };
var afterLinkFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var afterLinkGameCenter = function (ctx, logger, nk, data, request) { };
var afterLinkGoogle = function (ctx, logger, nk, data, request) { };
var afterLinkSteam = function (ctx, logger, nk, data, request) { };
var afterListChannelMessages = function (ctx, logger, nk, data, request) { };
var afterListFriends = function (ctx, logger, nk, data, request) { };
var afterListGroupUsers = function (ctx, logger, nk, data, request) { };
var afterListGroups = function (ctx, logger, nk, data, request) { };
var afterListLeaderboardRecords = function (ctx, logger, nk, data, request) { };
var afterListLeaderboardRecordsAroundOwner = function (ctx, logger, nk, data, request) { };
var afterListMatches = function (ctx, logger, nk, data, request) { };
var afterListNotifications = function (ctx, logger, nk, data, request) { };
var afterListStorageObjects = function (ctx, logger, nk, data, request) { };
var afterListTournamentRecords = function (ctx, logger, nk, data, request) { };
var afterListTournamentRecordsAroundOwner = function (ctx, logger, nk, data, request) { };
var afterListTournaments = function (ctx, logger, nk, data, request) { };
var afterListUserGroups = function (ctx, logger, nk, data, request) { };
var afterMatchCreate = function (ctx, logger, nk, output, input) {
    if (typeof input.matchCreate !==
        "undefined" &&
        output !== null &&
        typeof output.matchCreate !==
            "undefined") {
    }
};
var afterMatchDataSend = function (ctx, logger, nk, output, input) {
    if (typeof input.matchDataSend !==
        "undefined" &&
        output !== null &&
        typeof output.matchDataSend !==
            "undefined") {
    }
};
var afterMatchJoin = function (ctx, logger, nk, output, input) {
    if (typeof input.matchJoin !== "undefined" &&
        output !== null &&
        typeof output.matchJoin !== "undefined") {
    }
};
var afterMatchLeave = function (ctx, logger, nk, output, input) {
    if (typeof input.matchLeave !== "undefined" &&
        output !== null &&
        typeof output.matchLeave !== "undefined") {
    }
};
var afterMatchmakerAdd = function (ctx, logger, nk, output, input) {
    if (typeof input.matchmakerAdd !==
        "undefined" &&
        output !== null &&
        typeof output.matchmakerAdd !==
            "undefined") {
    }
};
var afterMatchmakerRemove = function (ctx, logger, nk, output, input) {
    if (typeof input.matchmakerRemove !==
        "undefined" &&
        output !== null &&
        typeof output.matchmakerRemove !==
            "undefined") {
    }
};
var afterPing = function (ctx, logger, nk, output, input) {
    if (typeof input.ping !== "undefined" &&
        output !== null &&
        typeof output.ping !== "undefined") {
    }
};
var afterPong = function (ctx, logger, nk, output, input) {
    if (typeof input.pong !== "undefined" &&
        output !== null &&
        typeof output.pong !== "undefined") {
    }
};
var afterPromoteGroupUsers = function (ctx, logger, nk, data, request) { };
var afterReadStorageObjects = function (ctx, logger, nk, data, request) { };
var afterStatusFollow = function (ctx, logger, nk, output, input) {
    if (typeof input.statusFollow !==
        "undefined" &&
        output !== null &&
        typeof output.statusFollow !==
            "undefined") {
    }
};
var afterStatusUnfollow = function (ctx, logger, nk, output, input) {
    if (typeof input.statusUnfollow !==
        "undefined" &&
        output !== null &&
        typeof output.statusUnfollow !==
            "undefined") {
    }
};
var afterStatusUpdate = function (ctx, logger, nk, output, input) {
    if (typeof input.statusUpdate !==
        "undefined" &&
        output !== null &&
        typeof output.statusUpdate !==
            "undefined") {
    }
};
var afterUnlinkApple = function (ctx, logger, nk, data, request) { };
var afterUnlinkCustom = function (ctx, logger, nk, data, request) { };
var afterUnlinkDevice = function (ctx, logger, nk, data, request) { };
var afterUnlinkEmail = function (ctx, logger, nk, data, request) { };
var afterUnlinkFacebook = function (ctx, logger, nk, data, request) { };
var afterUnlinkFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var afterUnlinkGameCenter = function (ctx, logger, nk, data, request) { };
var afterUnlinkGoogle = function (ctx, logger, nk, data, request) { };
var afterUnlinkSteam = function (ctx, logger, nk, data, request) { };
var afterUpdateAccount = function (ctx, logger, nk, data, request) { };
var afterUpdateGroup = function (ctx, logger, nk, data, request) { };
var afterValidatePurchaseApple = function (ctx, logger, nk, data, request) { };
var afterValidatePurchaseGoogle = function (ctx, logger, nk, data, request) { };
var afterValidatePurchaseHuawei = function (ctx, logger, nk, data, request) { };
var afterWriteLeaderboardRecord = function (ctx, logger, nk, data, request) { };
var afterWriteStorageObjects = function (ctx, logger, nk, data, request) { };
var afterWriteTournamentRecord = function (ctx, logger, nk, data, request) { };
var beforeAddFriends = function (ctx, logger, nk, data, request) { };
var beforeAddGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateApple = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateCustom = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateDevice = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateEmail = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateFacebook = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateGameCenter = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateGoogle = function (ctx, logger, nk, data, request) { };
var beforeAuthenticateSteam = function (ctx, logger, nk, data, request) { };
var beforeBanGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeBlockFriends = function (ctx, logger, nk, data, request) { };
var beforeChannelJoin = function (ctx, logger, nk, envlope) {
    if (typeof envlope.channelJoin !==
        "undefined") {
    }
};
var beforeChannelLeave = function (ctx, logger, nk, envlope) {
    if (typeof envlope.channelLeave !==
        "undefined") {
    }
};
var beforeChannelMessageRemove = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageRemove !== "undefined") {
    }
};
var beforeChannelMessageSend = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageSend !== "undefined") {
    }
};
var beforeChannelMessageUpdate = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageUpdate !== "undefined") {
    }
};
var beforeCreateGroup = function (ctx, logger, nk, data, request) { };
var beforeDeleteFriends = function (ctx, logger, nk, data, request) { };
var beforeDeleteGroup = function (ctx, logger, nk, data, request) { };
var beforeDeleteLeaderboardRecord = function (ctx, logger, nk, data, request) { };
var beforeDeleteNotifications = function (ctx, logger, nk, data, request) { };
var beforeDeleteStorageObjects = function (ctx, logger, nk, data, request) { };
var beforeDemoteGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeEvent = function (ctx, logger, nk, data, request) { };
var beforeGetAccount = function (ctx, logger, nk, data, request) { };
var beforeGetUsers = function (ctx, logger, nk, data, request) { };
var beforeImportFacebookFriends = function (ctx, logger, nk, data, request) { };
var beforeImportSteamFriends = function (ctx, logger, nk, data, request) { };
var beforeJoinGroup = function (ctx, logger, nk, data, request) { };
var beforeJoinTournament = function (ctx, logger, nk, data, request) { };
var beforeKickGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeLeaveGroup = function (ctx, logger, nk, data, request) { };
var beforeLinkApple = function (ctx, logger, nk, data, request) { };
var beforeLinkCustom = function (ctx, logger, nk, data, request) { };
var beforeLinkDevice = function (ctx, logger, nk, data, request) { };
var beforeLinkEmail = function (ctx, logger, nk, data, request) { };
var beforeLinkFacebook = function (ctx, logger, nk, data, request) { };
var beforeLinkFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var beforeLinkGameCenter = function (ctx, logger, nk, data, request) { };
var beforeLinkGoogle = function (ctx, logger, nk, data, request) { };
var beforeLinkSteam = function (ctx, logger, nk, data, request) { };
var beforeListChannelMessages = function (ctx, logger, nk, data, request) { };
var beforeListFriends = function (ctx, logger, nk, data, request) { };
var beforeListGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeListGroups = function (ctx, logger, nk, data, request) { };
var beforeListLeaderboardRecords = function (ctx, logger, nk, data, request) { };
var beforeListLeaderboardRecordsAroundOwner = function (ctx, logger, nk, data, request) { };
var beforeListMatches = function (ctx, logger, nk, data, request) { };
var beforeListNotifications = function (ctx, logger, nk, data, request) { };
var beforeListStorageObjects = function (ctx, logger, nk, data, request) { };
var beforeListTournamentRecords = function (ctx, logger, nk, data, request) { };
var beforeListTournamentRecordsAroundOwner = function (ctx, logger, nk, data, request) { };
var beforeListTournaments = function (ctx, logger, nk, data, request) { };
var beforeListUserGroups = function (ctx, logger, nk, data, request) { };
var beforeMatchCreate = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchCreate !==
        "undefined") {
    }
};
var beforeMatchDataSend = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchDataSend !==
        "undefined") {
    }
};
var beforeMatchJoin = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchJoin !== "undefined") {
    }
};
var beforeMatchLeave = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchLeave !== "undefined") {
    }
};
var beforeMatchmakerAdd = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchmakerAdd !==
        "undefined") {
    }
};
var beforeMatchmakerRemove = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchmakerRemove !==
        "undefined") {
    }
};
var beforePing = function (ctx, logger, nk, envlope) {
    if (typeof envlope.ping !== "undefined") {
    }
};
var beforePong = function (ctx, logger, nk, envlope) {
    if (typeof envlope.pong !== "undefined") {
    }
};
var beforePromoteGroupUsers = function (ctx, logger, nk, data, request) { };
var beforeReadStorageObjects = function (ctx, logger, nk, data, request) { };
var beforeStatusFollow = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusFollow !==
        "undefined") {
    }
};
var beforeStatusUnfollow = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusUnfollow !==
        "undefined") {
    }
};
var beforeStatusUpdate = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusUpdate !==
        "undefined") {
    }
};
var beforeUnlinkApple = function (ctx, logger, nk, data, request) { };
var beforeUnlinkCustom = function (ctx, logger, nk, data, request) { };
var beforeUnlinkDevice = function (ctx, logger, nk, data, request) { };
var beforeUnlinkEmail = function (ctx, logger, nk, data, request) { };
var beforeUnlinkFacebook = function (ctx, logger, nk, data, request) { };
var beforeUnlinkFacebookInstantGame = function (ctx, logger, nk, data, request) { };
var beforeUnlinkGameCenter = function (ctx, logger, nk, data, request) { };
var beforeUnlinkGoogle = function (ctx, logger, nk, data, request) { };
var beforeUnlinkSteam = function (ctx, logger, nk, data, request) { };
var beforeUpdateAccount = function (ctx, logger, nk, data, request) { };
var beforeUpdateGroup = function (ctx, logger, nk, data, request) { };
var beforeValidatePurchaseApple = function (ctx, logger, nk, data, request) { };
var beforeValidatePurchaseGoogle = function (ctx, logger, nk, data, request) { };
var beforeValidatePurchaseHuawei = function (ctx, logger, nk, data, request) { };
var beforeWriteLeaderboardRecord = function (ctx, logger, nk, data, request) { };
var beforeWriteStorageObjects = function (ctx, logger, nk, data, request) { };
var beforeWriteTournamentRecord = function (ctx, logger, nk, data, request) { };
var InitModule = function (_ctx, logger, nk, initializer) {
    logger.info("InitModule called");
    initializer.registerAfterAddFriends(afterAddFriends);
    initializer.registerAfterAddGroupUsers(afterAddGroupUsers);
    initializer.registerAfterAuthenticateApple(afterAuthenticateApple);
    initializer.registerAfterAuthenticateCustom(afterAuthenticateCustom);
    initializer.registerAfterAuthenticateDevice(afterAuthenticateDevice);
    initializer.registerAfterAuthenticateEmail(afterAuthenticateEmail);
    initializer.registerAfterAuthenticateFacebook(afterAuthenticateFacebook);
    initializer.registerAfterAuthenticateFacebookInstantGame(afterAuthenticateFacebookInstantGame);
    initializer.registerAfterAuthenticateGameCenter(afterAuthenticateGameCenter);
    initializer.registerAfterAuthenticateGoogle(afterAuthenticateGoogle);
    initializer.registerAfterAuthenticateSteam(afterAuthenticateSteam);
    initializer.registerAfterBanGroupUsers(afterBanGroupUsers);
    initializer.registerAfterBlockFriends(afterBlockFriends);
    initializer.registerAfterCreateGroup(afterCreateGroup);
    initializer.registerAfterDeleteFriends(afterDeleteFriends);
    initializer.registerAfterDeleteGroup(afterDeleteGroup);
    initializer.registerAfterDeleteLeaderboardRecord(afterDeleteLeaderboardRecord);
    initializer.registerAfterDeleteNotifications(afterDeleteNotifications);
    initializer.registerAfterDeleteStorageObjects(afterDeleteStorageObjects);
    initializer.registerAfterDemoteGroupUsers(afterDemoteGroupUsers);
    initializer.registerAfterEvent(afterEvent);
    initializer.registerAfterGetAccount(afterGetAccount);
    initializer.registerAfterGetUsers(afterGetUsers);
    initializer.registerAfterImportFacebookFriends(afterImportFacebookFriends);
    initializer.registerAfterImportSteamFriends(afterImportSteamFriends);
    initializer.registerAfterJoinGroup(afterJoinGroup);
    initializer.registerAfterJoinTournament(afterJoinTournament);
    initializer.registerAfterKickGroupUsers(afterKickGroupUsers);
    initializer.registerAfterLeaveGroup(afterLeaveGroup);
    initializer.registerAfterLinkApple(afterLinkApple);
    initializer.registerAfterLinkCustom(afterLinkCustom);
    initializer.registerAfterLinkDevice(afterLinkDevice);
    initializer.registerAfterLinkEmail(afterLinkEmail);
    initializer.registerAfterLinkFacebook(afterLinkFacebook);
    initializer.registerAfterLinkFacebookInstantGame(afterLinkFacebookInstantGame);
    initializer.registerAfterLinkGameCenter(afterLinkGameCenter);
    initializer.registerAfterLinkGoogle(afterLinkGoogle);
    initializer.registerAfterLinkSteam(afterLinkSteam);
    initializer.registerAfterListChannelMessages(afterListChannelMessages);
    initializer.registerAfterListFriends(afterListFriends);
    initializer.registerAfterListGroupUsers(afterListGroupUsers);
    initializer.registerAfterListGroups(afterListGroups);
    initializer.registerAfterListLeaderboardRecords(afterListLeaderboardRecords);
    initializer.registerAfterListLeaderboardRecordsAroundOwner(afterListLeaderboardRecordsAroundOwner);
    initializer.registerAfterListMatches(afterListMatches);
    initializer.registerAfterListNotifications(afterListNotifications);
    initializer.registerAfterListStorageObjects(afterListStorageObjects);
    initializer.registerAfterListTournamentRecords(afterListTournamentRecords);
    initializer.registerAfterListTournamentRecordsAroundOwner(afterListTournamentRecordsAroundOwner);
    initializer.registerAfterListTournaments(afterListTournaments);
    initializer.registerAfterListUserGroups(afterListUserGroups);
    initializer.registerAfterPromoteGroupUsers(afterPromoteGroupUsers);
    initializer.registerAfterReadStorageObjects(afterReadStorageObjects);
    initializer.registerAfterUnlinkApple(afterUnlinkApple);
    initializer.registerAfterUnlinkCustom(afterUnlinkCustom);
    initializer.registerAfterUnlinkDevice(afterUnlinkDevice);
    initializer.registerAfterUnlinkEmail(afterUnlinkEmail);
    initializer.registerAfterUnlinkFacebook(afterUnlinkFacebook);
    initializer.registerAfterUnlinkFacebookInstantGame(afterUnlinkFacebookInstantGame);
    initializer.registerAfterUnlinkGameCenter(afterUnlinkGameCenter);
    initializer.registerAfterUnlinkGoogle(afterUnlinkGoogle);
    initializer.registerAfterUnlinkSteam(afterUnlinkSteam);
    initializer.registerAfterUpdateAccount(afterUpdateAccount);
    initializer.registerAfterUpdateGroup(afterUpdateGroup);
    initializer.registerAfterValidatePurchaseApple(afterValidatePurchaseApple);
    initializer.registerAfterValidatePurchaseGoogle(afterValidatePurchaseGoogle);
    initializer.registerAfterValidatePurchaseHuawei(afterValidatePurchaseHuawei);
    initializer.registerAfterWriteLeaderboardRecord(afterWriteLeaderboardRecord);
    initializer.registerAfterWriteStorageObjects(afterWriteStorageObjects);
    initializer.registerAfterWriteTournamentRecord(afterWriteTournamentRecord);
    initializer.registerBeforeAddFriends(beforeAddFriends);
    initializer.registerBeforeAddGroupUsers(beforeAddGroupUsers);
    initializer.registerBeforeAuthenticateApple(beforeAuthenticateApple);
    initializer.registerBeforeAuthenticateCustom(beforeAuthenticateCustom);
    initializer.registerBeforeAuthenticateDevice(beforeAuthenticateDevice);
    initializer.registerBeforeAuthenticateEmail(beforeAuthenticateEmail);
    initializer.registerBeforeAuthenticateFacebook(beforeAuthenticateFacebook);
    initializer.registerBeforeAuthenticateFacebookInstantGame(beforeAuthenticateFacebookInstantGame);
    initializer.registerBeforeAuthenticateGameCenter(beforeAuthenticateGameCenter);
    initializer.registerBeforeAuthenticateGoogle(beforeAuthenticateGoogle);
    initializer.registerBeforeAuthenticateSteam(beforeAuthenticateSteam);
    initializer.registerBeforeBanGroupUsers(beforeBanGroupUsers);
    initializer.registerBeforeBlockFriends(beforeBlockFriends);
    initializer.registerBeforeCreateGroup(beforeCreateGroup);
    initializer.registerBeforeDeleteFriends(beforeDeleteFriends);
    initializer.registerBeforeDeleteGroup(beforeDeleteGroup);
    initializer.registerBeforeDeleteLeaderboardRecord(beforeDeleteLeaderboardRecord);
    initializer.registerBeforeDeleteNotifications(beforeDeleteNotifications);
    initializer.registerBeforeDeleteStorageObjects(beforeDeleteStorageObjects);
    initializer.registerBeforeDemoteGroupUsers(beforeDemoteGroupUsers);
    initializer.registerBeforeEvent(beforeEvent);
    initializer.registerBeforeGetAccount(beforeGetAccount);
    initializer.registerBeforeGetUsers(beforeGetUsers);
    initializer.registerBeforeImportFacebookFriends(beforeImportFacebookFriends);
    initializer.registerBeforeImportSteamFriends(beforeImportSteamFriends);
    initializer.registerBeforeJoinGroup(beforeJoinGroup);
    initializer.registerBeforeJoinTournament(beforeJoinTournament);
    initializer.registerBeforeKickGroupUsers(beforeKickGroupUsers);
    initializer.registerBeforeLeaveGroup(beforeLeaveGroup);
    initializer.registerBeforeLinkApple(beforeLinkApple);
    initializer.registerBeforeLinkCustom(beforeLinkCustom);
    initializer.registerBeforeLinkDevice(beforeLinkDevice);
    initializer.registerBeforeLinkEmail(beforeLinkEmail);
    initializer.registerBeforeLinkFacebook(beforeLinkFacebook);
    initializer.registerBeforeLinkFacebookInstantGame(beforeLinkFacebookInstantGame);
    initializer.registerBeforeLinkGameCenter(beforeLinkGameCenter);
    initializer.registerBeforeLinkGoogle(beforeLinkGoogle);
    initializer.registerBeforeLinkSteam(beforeLinkSteam);
    initializer.registerBeforeListChannelMessages(beforeListChannelMessages);
    initializer.registerBeforeListFriends(beforeListFriends);
    initializer.registerBeforeListGroupUsers(beforeListGroupUsers);
    initializer.registerBeforeListGroups(beforeListGroups);
    initializer.registerBeforeListLeaderboardRecords(beforeListLeaderboardRecords);
    initializer.registerBeforeListLeaderboardRecordsAroundOwner(beforeListLeaderboardRecordsAroundOwner);
    initializer.registerBeforeListMatches(beforeListMatches);
    initializer.registerBeforeListNotifications(beforeListNotifications);
    initializer.registerBeforeListStorageObjects(beforeListStorageObjects);
    initializer.registerBeforeListTournamentRecords(beforeListTournamentRecords);
    initializer.registerBeforeListTournamentRecordsAroundOwner(beforeListTournamentRecordsAroundOwner);
    initializer.registerBeforeListTournaments(beforeListTournaments);
    initializer.registerBeforeListUserGroups(beforeListUserGroups);
    initializer.registerBeforePromoteGroupUsers(beforePromoteGroupUsers);
    initializer.registerBeforeReadStorageObjects(beforeReadStorageObjects);
    initializer.registerBeforeUnlinkApple(beforeUnlinkApple);
    initializer.registerBeforeUnlinkCustom(beforeUnlinkCustom);
    initializer.registerBeforeUnlinkDevice(beforeUnlinkDevice);
    initializer.registerBeforeUnlinkEmail(beforeUnlinkEmail);
    initializer.registerBeforeUnlinkFacebook(beforeUnlinkFacebook);
    initializer.registerBeforeUnlinkFacebookInstantGame(beforeUnlinkFacebookInstantGame);
    initializer.registerBeforeUnlinkGameCenter(beforeUnlinkGameCenter);
    initializer.registerBeforeUnlinkGoogle(beforeUnlinkGoogle);
    initializer.registerBeforeUnlinkSteam(beforeUnlinkSteam);
    initializer.registerBeforeUpdateAccount(beforeUpdateAccount);
    initializer.registerBeforeUpdateGroup(beforeUpdateGroup);
    initializer.registerBeforeValidatePurchaseApple(beforeValidatePurchaseApple);
    initializer.registerBeforeValidatePurchaseGoogle(beforeValidatePurchaseGoogle);
    initializer.registerBeforeValidatePurchaseHuawei(beforeValidatePurchaseHuawei);
    initializer.registerBeforeWriteLeaderboardRecord(beforeWriteLeaderboardRecord);
    initializer.registerBeforeWriteStorageObjects(beforeWriteStorageObjects);
    initializer.registerBeforeWriteTournamentRecord(beforeWriteTournamentRecord);
    initializer.registerRtAfter("ChannelJoin", afterChannelJoin);
    initializer.registerRtAfter("ChannelLeave", afterChannelLeave);
    initializer.registerRtAfter("ChannelMessageRemove", afterChannelMessageRemove);
    initializer.registerRtAfter("ChannelMessageSend", afterChannelMessageSend);
    initializer.registerRtAfter("ChannelMessageUpdate", afterChannelMessageUpdate);
    initializer.registerRtAfter("MatchCreate", afterMatchCreate);
    initializer.registerRtAfter("MatchDataSend", afterMatchDataSend);
    initializer.registerRtAfter("MatchJoin", afterMatchJoin);
    initializer.registerRtAfter("MatchLeave", afterMatchLeave);
    initializer.registerRtAfter("MatchmakerAdd", afterMatchmakerAdd);
    initializer.registerRtAfter("MatchmakerRemove", afterMatchmakerRemove);
    initializer.registerRtAfter("Ping", afterPing);
    initializer.registerRtAfter("Pong", afterPong);
    initializer.registerRtAfter("StatusFollow", afterStatusFollow);
    initializer.registerRtAfter("StatusUnfollow", afterStatusUnfollow);
    initializer.registerRtAfter("StatusUpdate", afterStatusUpdate);
    initializer.registerRtBefore("ChannelJoin", beforeChannelJoin);
    initializer.registerRtBefore("ChannelLeave", beforeChannelLeave);
    initializer.registerRtBefore("ChannelMessageRemove", beforeChannelMessageRemove);
    initializer.registerRtBefore("ChannelMessageSend", beforeChannelMessageSend);
    initializer.registerRtBefore("ChannelMessageUpdate", beforeChannelMessageUpdate);
    initializer.registerRtBefore("MatchCreate", beforeMatchCreate);
    initializer.registerRtBefore("MatchDataSend", beforeMatchDataSend);
    initializer.registerRtBefore("MatchJoin", beforeMatchJoin);
    initializer.registerRtBefore("MatchLeave", beforeMatchLeave);
    initializer.registerRtBefore("MatchmakerAdd", beforeMatchmakerAdd);
    initializer.registerRtBefore("MatchmakerRemove", beforeMatchmakerRemove);
    initializer.registerRtBefore("Ping", beforePing);
    initializer.registerRtBefore("Pong", beforePong);
    initializer.registerRtBefore("StatusFollow", beforeStatusFollow);
    initializer.registerRtBefore("StatusUnfollow", beforeStatusUnfollow);
    initializer.registerRtBefore("StatusUpdate", beforeStatusUpdate);
    initializer.registerLeaderboardReset(registerLeaderboardReset);
    //registerMatch < State = MatchState > (name: string, functions: MatchHandler<State>): void;
    initializer.registerMatch("", match);
    initializer.registerMatchmakerMatched(matchmakerMatched);
    initializer.registerTournamentEnd(tournamentEnd);
    initializer.registerTournamentReset(tournamentReset);
    initializer.registerRpc("healthcheck", rpcHealthCheck);
    nk.httpRequest("http://redpanda:8082/topics/nakama", "post", {
        "Content-Type": "application/vnd.kafka.json.v2+json",
    }, JSON.stringify({
        records: [{
                key: "nakama", value: {
                    message: "hi", user: "nakama"
                }
            }]
    }));
};
var registerLeaderboardReset = function (ctx, logger, nk, leaderboard, reset) { };
// registerMatch < State = MatchState > (name: string, functions: MatchHandler<State>): void;
var matchInit = function (ctx, logger, nk, params) {
    return { state: {}, tickRate: 0, label: "" };
};
var matchJoinAttempt = function (ctx, logger, nk, dispatcher, tick, state, presence, metadata) {
    return { state: {}, accept: false, rejectMessage: "" };
};
var matchJoin = function (ctx, logger, nk, dispatcher, tick, state, presences) {
    return { state: {} };
};
var matchLeave = function (ctx, logger, nk, dispatcher, tick, state, presences) {
    return { state: {} };
};
var matchLoop = function (ctx, logger, nk, dispatcher, tick, state, messages) {
    return { state: {} };
};
var matchTerminate = function (ctx, logger, nk, dispatcher, tick, state, graceSeconds) {
    return { state: {} };
};
var matchSignal = function (ctx, logger, nk, dispatcher, tick, state, data) {
    return { state: {}, date: String(new Date()) };
};
var match = {
    matchInit: matchInit,
    matchJoinAttempt: matchJoinAttempt,
    matchJoin: matchJoin,
    matchLeave: matchLeave,
    matchLoop: matchLoop,
    matchTerminate: matchTerminate,
    matchSignal: matchSignal
};
var matchmakerMatched = function (ctx, logger, nk, matches) { };
var tournamentEnd = function (ctx, logger, nk, tournament, end, reset) { };
var tournamentReset = function (ctx, logger, nk, tournament, end, reset) { };
var rpcHealthCheck = function (_ctx, logger, _nk, _initializer) {
    logger.info("rpcHealthCheck called");
    return JSON.stringify({ success: true });
};
