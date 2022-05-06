"use strict";
var redpanda = function (ctx, logger, nk, payload) {
    nk.httpRequest("http://redpanda:8082/topics/nakama", "post", {
        "Content-Type": "application/vnd.kafka.json.v2+json"
    }, JSON.stringify({
        "records": [{
                "key": { "node": ctx.node },
                "value": { ctx: ctx, payload: payload }
            }]
    }));
};
var afterAddFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAddFriends", data: data, request: request });
};
var afterAddGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAddGroupUsers", data: data, request: request });
};
var afterAuthenticateApple = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateApple", data: data, request: request });
};
var afterAuthenticateCustom = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateCustom", data: data, request: request });
};
var afterAuthenticateDevice = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateDevice", data: data, request: request });
};
var afterAuthenticateEmail = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateEmail", data: data, request: request });
};
var afterAuthenticateFacebook = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateFacebook", data: data, request: request });
};
var afterAuthenticateFacebookInstantGame = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateFacebookInstantGame", data: data, request: request });
};
var afterAuthenticateGameCenter = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateGameCenter", data: data, request: request });
};
var afterAuthenticateGoogle = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateGoogle", data: data, request: request });
};
var afterAuthenticateSteam = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterAuthenticateSteam", data: data, request: request });
};
var afterBanGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterBanGroupUsers", data: data, request: request });
};
var afterBlockFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterBlockFriends", data: data, request: request });
};
var afterChannelJoin = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterChannelJoin", output: output, input: input });
    if (typeof input.channelJoin !==
        "undefined" &&
        output !== null &&
        typeof output.channelJoin !== "undefined") {
    }
    return output;
};
var afterChannelLeave = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterChannelLeave", output: output, input: input });
    if (typeof input.channelLeave !==
        "undefined" &&
        output !== null &&
        typeof output.channelLeave !==
            "undefined") {
    }
    return output;
};
var afterChannelMessageRemove = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterChannelMessageRemove", output: output, input: input });
    if (typeof input
        .channelMessageRemove !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageRemove !== "undefined") {
    }
    return output;
};
var afterChannelMessageSend = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterChannelMessageSend", output: output, input: input });
    if (typeof input
        .channelMessageSend !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageSend !== "undefined") {
    }
    return output;
};
var afterChannelMessageUpdate = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterChannelMessageUpdate", output: output, input: input });
    if (typeof input
        .channelMessageUpdate !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageUpdate !== "undefined") {
    }
    return output;
};
var afterCreateGroup = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterCreateGroup", data: data, request: request });
};
var afterDeleteFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDeleteFriends", data: data, request: request });
};
var afterDeleteGroup = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDeleteGroup", data: data, request: request });
};
var afterDeleteLeaderboardRecord = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDeleteLeaderboardRecord", data: data, request: request });
};
var afterDeleteNotification = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDeleteNotification", data: data, request: request });
};
var afterDeleteStorageObjects = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDeleteStorageObjects", data: data, request: request });
};
var afterDemoteGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterDemoteGroupUsers", data: data, request: request });
};
var afterEvent = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterEvent", data: data, request: request });
};
var afterGetAccount = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterGetAccount", data: data, request: request });
};
var afterGetUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterGetUsers", data: data, request: request });
};
var afterImportFacebookFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterImportFacebookFriends", data: data, request: request });
};
var afterImportSteamFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterImportSteamFriends", data: data, request: request });
};
var afterJoinGroup = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterJoinGroup", data: data, request: request });
};
var afterJoinTournament = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterJoinTournament", data: data, request: request });
};
var afterKickGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterKickGroupUsers", data: data, request: request });
};
var afterLeaveGroup = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLeaveGroup", data: data, request: request });
};
var afterLinkApple = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkApple", data: data, request: request });
};
var afterLinkCustom = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkCustom", data: data, request: request });
};
var afterLinkDevice = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkDevice", data: data, request: request });
};
var afterLinkEmail = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkEmail", data: data, request: request });
};
var afterLinkFacebook = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkFacebook", data: data, request: request });
};
var afterLinkFacebookInstantGame = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkFacebookInstantGame", data: data, request: request });
};
var afterLinkGameCenter = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkGameCenter", data: data, request: request });
};
var afterLinkGoogle = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkGoogle", data: data, request: request });
};
var afterLinkSteam = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterLinkSteam", data: data, request: request });
};
var afterListChannelMessages = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListChannelMessages", data: data, request: request });
};
var afterListFriends = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListFriends", data: data, request: request });
};
var afterListGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListGroupUsers", data: data, request: request });
};
var afterListGroups = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListGroups", data: data, request: request });
};
var afterListLeaderboardRecords = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListLeaderboardRecords", data: data, request: request });
};
var afterListLeaderboardRecordsAroundOwner = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListLeaderboardRecordsAroundOwner", data: data, request: request });
};
var afterListMatches = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListMatches", data: data, request: request });
};
var afterListNotifications = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListNotifications", data: data, request: request });
};
var afterListStorageObjects = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListStorageObjects", data: data, request: request });
};
var afterListTournamentRecords = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListTournamentRecords", data: data, request: request });
};
var afterListTournamentRecordsAroundOwner = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListTournamentRecordsAroundOwner", data: data, request: request });
};
var afterListTournaments = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListTournaments", data: data, request: request });
};
var afterListUserGroups = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterListUserGroups", data: data, request: request });
};
var afterMatchCreate = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchCreate", output: output, input: input });
    if (typeof input.matchCreate !==
        "undefined" &&
        output !== null &&
        typeof output.matchCreate !==
            "undefined") {
    }
    return output;
};
var afterMatchDataSend = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchDataSend", output: output, input: input });
    if (typeof input.matchDataSend !==
        "undefined" &&
        output !== null &&
        typeof output.matchDataSend !==
            "undefined") {
    }
    return output;
};
var afterMatchJoin = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchJoin", output: output, input: input });
    if (typeof input.matchJoin !== "undefined" &&
        output !== null &&
        typeof output.matchJoin !== "undefined") {
    }
    return output;
};
var afterMatchLeave = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchLeave", output: output, input: input });
    if (typeof input.matchLeave !== "undefined" &&
        output !== null &&
        typeof output.matchLeave !== "undefined") {
    }
    return output;
};
var afterMatchmakerAdd = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchmakerAdd", output: output, input: input });
    if (typeof input.matchmakerAdd !==
        "undefined" &&
        output !== null &&
        typeof output.matchmakerAdd !==
            "undefined") {
    }
    return output;
};
var afterMatchmakerRemove = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterMatchmakerRemove", output: output, input: input });
    if (typeof input.matchmakerRemove !==
        "undefined" &&
        output !== null &&
        typeof output.matchmakerRemove !==
            "undefined") {
    }
    return output;
};
var afterPing = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterPing", output: output, input: input });
    if (typeof input.ping !== "undefined" &&
        output !== null &&
        typeof output.ping !== "undefined") {
    }
    return output;
};
var afterPong = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterPong", output: output, input: input });
    if (typeof input.pong !== "undefined" &&
        output !== null &&
        typeof output.pong !== "undefined") {
    }
    return output;
};
var afterPromoteGroupUsers = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterPromoteGroupUsers", data: data, request: request });
};
var afterReadStorageObjects = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterReadStorageObjects", data: data, request: request });
};
var afterStatusFollow = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterStatusFollow", output: output, input: input });
    if (typeof input.statusFollow !==
        "undefined" &&
        output !== null &&
        typeof output.statusFollow !==
            "undefined") {
    }
    return output;
};
var afterStatusUnfollow = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterStatusUnfollow", output: output, input: input });
    if (typeof input.statusUnfollow !==
        "undefined" &&
        output !== null &&
        typeof output.statusUnfollow !==
            "undefined") {
    }
    return output;
};
var afterStatusUpdate = function (ctx, logger, nk, output, input) {
    redpanda(ctx, logger, nk, { name: "afterStatusUpdate", output: output, input: input });
    if (typeof input.statusUpdate !==
        "undefined" &&
        output !== null &&
        typeof output.statusUpdate !==
            "undefined") {
    }
    return output;
};
var afterUnlinkApple = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkApple", data: data, request: request });
};
var afterUnlinkCustom = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkCustom", data: data, request: request });
};
var afterUnlinkDevice = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkDevice", data: data, request: request });
};
var afterUnlinkEmail = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkEmail", data: data, request: request });
};
var afterUnlinkFacebook = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkFacebook", data: data, request: request });
};
var afterUnlinkFacebookInstantGame = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkFacebookInstantGame", data: data, request: request });
};
var afterUnlinkGameCenter = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkGameCenter", data: data, request: request });
};
var afterUnlinkGoogle = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkGoogle", data: data, request: request });
};
var afterUnlinkSteam = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUnlinkSteam", data: data, request: request });
};
var afterUpdateAccount = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUpdateAccount", data: data, request: request });
};
var afterUpdateGroup = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterUpdateGroup", data: data, request: request });
};
var afterValidatePurchaseApple = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterValidatePurchaseApple", data: data, request: request });
};
var afterValidatePurchaseGoogle = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterValidatePurchaseGoogle", data: data, request: request });
};
var afterValidatePurchaseHuawei = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterValidatePurchaseHuawei", data: data, request: request });
};
var afterWriteLeaderboardRecord = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterWriteLeaderboardRecord", data: data, request: request });
};
var afterWriteStorageObjects = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterWriteStorageObjects", data: data, request: request });
};
var afterWriteTournamentRecord = function (ctx, logger, nk, data, request) {
    redpanda(ctx, logger, nk, { name: "afterWriteTournamentRecord", data: data, request: request });
};
var beforeAddFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeAddGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateApple = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateCustom = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateDevice = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateEmail = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateFacebook = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateFacebookInstantGame = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateGameCenter = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateGoogle = function (ctx, logger, nk, data) {
    return data;
};
var beforeAuthenticateSteam = function (ctx, logger, nk, data) {
    return data;
};
var beforeBanGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeBlockFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeChannelJoin = function (ctx, logger, nk, envlope) {
    if (typeof envlope.channelJoin !==
        "undefined") {
    }
    return envlope;
};
var beforeChannelLeave = function (ctx, logger, nk, envlope) {
    if (typeof envlope.channelLeave !==
        "undefined") {
    }
    return envlope;
};
var beforeChannelMessageRemove = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageRemove !== "undefined") {
    }
    return envlope;
};
var beforeChannelMessageSend = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageSend !== "undefined") {
    }
    return envlope;
};
var beforeChannelMessageUpdate = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageUpdate !== "undefined") {
    }
    return envlope;
};
var beforeCreateGroup = function (ctx, logger, nk, data) {
    return data;
};
var beforeDeleteFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeDeleteGroup = function (ctx, logger, nk, data) {
    return data;
};
var beforeDeleteLeaderboardRecord = function (ctx, logger, nk, data) {
    return data;
};
var beforeDeleteNotification = function (ctx, logger, nk, data) {
    return data;
};
var beforeDeleteStorageObjects = function (ctx, logger, nk, data) {
    return data;
};
var beforeDemoteGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeEvent = function (ctx, logger, nk, data) {
    return data;
};
var beforeGetAccount = function (ctx, logger, nk, data) {
    return data;
};
var beforeGetUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeImportFacebookFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeImportSteamFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeJoinGroup = function (ctx, logger, nk, data) {
    return data;
};
var beforeJoinTournament = function (ctx, logger, nk, data) {
    return data;
};
var beforeKickGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeLeaveGroup = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkApple = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkCustom = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkDevice = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkEmail = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkFacebook = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkFacebookInstantGame = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkGameCenter = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkGoogle = function (ctx, logger, nk, data) {
    return data;
};
var beforeLinkSteam = function (ctx, logger, nk, data) {
    return data;
};
var beforeListChannelMessages = function (ctx, logger, nk, data) {
    return data;
};
var beforeListFriends = function (ctx, logger, nk, data) {
    return data;
};
var beforeListGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeListGroups = function (ctx, logger, nk, data) {
    return data;
};
var beforeListLeaderboardRecords = function (ctx, logger, nk, data) {
    return data;
};
var beforeListLeaderboardRecordsAroundOwner = function (ctx, logger, nk, data) {
    return data;
};
var beforeListMatches = function (ctx, logger, nk, data) {
    return data;
};
var beforeListNotifications = function (ctx, logger, nk, data) {
    return data;
};
var beforeListStorageObjects = function (ctx, logger, nk, data) {
    return data;
};
var beforeListTournamentRecords = function (ctx, logger, nk, data) {
    return data;
};
var beforeListTournamentRecordsAroundOwner = function (ctx, logger, nk, data) {
    return data;
};
var beforeListTournaments = function (ctx, logger, nk, data) {
    return data;
};
var beforeListUserGroups = function (ctx, logger, nk, data) {
    return data;
};
var beforeMatchCreate = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchCreate !==
        "undefined") {
    }
    return envlope;
};
var beforeMatchDataSend = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchDataSend !==
        "undefined") {
    }
    return envlope;
};
var beforeMatchJoin = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchJoin !== "undefined") {
    }
    return envlope;
};
var beforeMatchLeave = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchLeave !== "undefined") {
    }
    return envlope;
};
var beforeMatchmakerAdd = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchmakerAdd !==
        "undefined") {
    }
    return envlope;
};
var beforeMatchmakerRemove = function (ctx, logger, nk, envlope) {
    if (typeof envlope.matchmakerRemove !==
        "undefined") {
    }
    return envlope;
};
var beforePing = function (ctx, logger, nk, envlope) {
    if (typeof envlope.ping !== "undefined") {
    }
};
var beforePong = function (ctx, logger, nk, envlope) {
    if (typeof envlope.pong !== "undefined") {
    }
};
var beforePromoteGroupUsers = function (ctx, logger, nk, data) {
    return data;
};
var beforeReadStorageObjects = function (ctx, logger, nk, data) {
    return data;
};
var beforeStatusFollow = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusFollow !==
        "undefined") {
    }
    return envlope;
};
var beforeStatusUnfollow = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusUnfollow !==
        "undefined") {
    }
    return envlope;
};
var beforeStatusUpdate = function (ctx, logger, nk, envlope) {
    if (typeof envlope.statusUpdate !==
        "undefined") {
    }
    return envlope;
};
var beforeUnlinkApple = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkCustom = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkDevice = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkEmail = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkFacebook = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkFacebookInstantGame = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkGameCenter = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkGoogle = function (ctx, logger, nk, data) {
    return data;
};
var beforeUnlinkSteam = function (ctx, logger, nk, data) {
    return data;
};
var beforeUpdateAccount = function (ctx, logger, nk, data) {
    return data;
};
var beforeUpdateGroup = function (ctx, logger, nk, data) {
    return data;
};
var beforeValidatePurchaseApple = function (ctx, logger, nk, data) {
    return data;
};
var beforeValidatePurchaseGoogle = function (ctx, logger, nk, data) {
    return data;
};
var beforeValidatePurchaseHuawei = function (ctx, logger, nk, data) {
    return data;
};
var beforeWriteLeaderboardRecord = function (ctx, logger, nk, data) {
    return data;
};
var beforeWriteStorageObjects = function (ctx, logger, nk, data) {
    return data;
};
var beforeWriteTournamentRecord = function (ctx, logger, nk, data) {
    return data;
};
var registerLeaderboardReset = function (ctx, logger, nk, leaderboard, reset) {
    redpanda(ctx, logger, nk, { name: "registerLeaderboardReset", leaderboard: leaderboard, reset: reset });
};
var matchInit = function (ctx, logger, nk, params) {
    redpanda(ctx, logger, nk, { name: "matchInit", params: params });
    return { state: {}, tickRate: 0, label: "" };
};
var matchJoinAttempt = function (ctx, logger, nk, dispatcher, tick, state, presence, metadata) {
    redpanda(ctx, logger, nk, { name: "matchJoinAttempt", tick: tick, state: state, presence: presence, metadata: metadata });
    return { state: {}, accept: false, rejectMessage: "" };
};
var matchJoin = function (ctx, logger, nk, dispatcher, tick, state, presences) {
    redpanda(ctx, logger, nk, { name: "matchJoin", tick: tick, state: state, presences: presences });
    return { state: {} };
};
var matchLeave = function (ctx, logger, nk, dispatcher, tick, state, presences) {
    redpanda(ctx, logger, nk, { name: "matchLeave", tick: tick, state: state, presences: presences });
    return { state: {} };
};
var matchLoop = function (ctx, logger, nk, dispatcher, tick, state, messages) {
    redpanda(ctx, logger, nk, { name: "matchLoop", tick: tick, state: state, messages: messages });
    return { state: {} };
};
var matchTerminate = function (ctx, logger, nk, dispatcher, tick, state, graceSeconds) {
    redpanda(ctx, logger, nk, { name: "matchTerminate", tick: tick, state: state, graceSeconds: graceSeconds });
    return { state: {} };
};
var matchSignal = function (ctx, logger, nk, dispatcher, tick, state, data) {
    redpanda(ctx, logger, nk, { name: "matchSignal", tick: tick, state: state, data: data });
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
var matchmakerMatched = function (ctx, logger, nk, matches) {
    redpanda(ctx, logger, nk, { name: "matchmakerMatched", matches: matches });
};
var tournamentEnd = function (ctx, logger, nk, tournament, end, reset) {
    redpanda(ctx, logger, nk, { name: "tournamentEnd", tournament: tournament, end: end, reset: reset });
};
var tournamentReset = function (ctx, logger, nk, tournament, end, reset) {
    redpanda(ctx, logger, nk, { name: "tournamentReset", tournament: tournament, end: end, reset: reset });
};
var InitModule = function (ctx, logger, nk, initializer) {
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
    initializer.registerAfterDeleteNotification(afterDeleteNotification);
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
    initializer.registerBeforeDeleteNotification(beforeDeleteNotification);
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
    initializer.registerMatch("", match);
    initializer.registerMatchmakerMatched(matchmakerMatched);
    initializer.registerTournamentEnd(tournamentEnd);
    initializer.registerTournamentReset(tournamentReset);
};
