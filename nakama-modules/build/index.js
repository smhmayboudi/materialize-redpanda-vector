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
var InitModule = function (ctx, logger, nk, initializer) {
    logger.info("InitModule called");
    initializer.registerRpc("healthcheck", rpcHealthCheck);
    var body3 = nk.httpRequest("http://redpanda:8082/topics/nakama", "post", {
        "Content-Type": "application/vnd.kafka.json.v2+json",
    }, JSON.stringify({
        "records": [{
                "key": {
                    "node": ctx.node
                },
                "value": {
                    clientIp: ctx.clientIp,
                    clientPort: ctx.clientPort,
                    env: ctx.env,
                    executionMode: ctx.executionMode,
                    headers: ctx.headers,
                    lang: ctx.lang,
                    matchId: ctx.matchId,
                    matchLabel: ctx.matchLabel,
                    matchNode: ctx.matchNode,
                    matchTickRate: ctx.matchTickRate,
                    node: ctx.node,
                    queryParams: ctx.queryParams,
                    sessionId: ctx.sessionId,
                    userId: ctx.userId,
                    userSessionExp: ctx.userSessionExp,
                    username: ctx.username,
                    vars: ctx.vars
                }
            }]
    }));
};
var registerLeaderboardReset = function (ctx, logger, nk, leaderboard, reset) { };
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
