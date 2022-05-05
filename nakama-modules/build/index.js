"use strict";
var InitModule = function (_ctx, logger, nk, initializer) {
    logger.info("InitModule called");
    initializer.registerRtAfter("ChannelJoin", afterChannelJoin);
    initializer.registerRtAfter("ChannelLeave", afterChannelLeave);
    initializer.registerRtAfter("ChannelMessageSend", afterChannelMessageSend);
    initializer.registerRtAfter("ChannelMessageUpdate", afterChannelMessageUpdate);
    initializer.registerRtAfter("ChannelMessageRemove", afterChannelMessageRemove);
    initializer.registerRtAfter("MatchCreate", afterMatchCreate);
    initializer.registerRtAfter("MatchDataSend", afterMatchDataSend);
    initializer.registerRtAfter("MatchJoin", afterMatchJoin);
    initializer.registerRtAfter("MatchLeave", afterMatchLeave);
    initializer.registerRtAfter("MatchmakerAdd", afterMatchmakerAdd);
    initializer.registerRtAfter("MatchmakerRemove", afterMatchmakerRemove);
    initializer.registerRtAfter("StatusFollow", afterStatusFollow);
    initializer.registerRtAfter("StatusUnfollow", afterStatusUnfollow);
    initializer.registerRtAfter("StatusUpdate", afterStatusUpdate);
    initializer.registerRtAfter("Ping", afterPing);
    initializer.registerRtAfter("Pong", afterPong);
    initializer.registerRtBefore("ChannelJoin", beforeChannelJoin);
    initializer.registerRtBefore("ChannelLeave", beforeChannelLeave);
    initializer.registerRtBefore("ChannelMessageSend", beforeChannelMessageSend);
    initializer.registerRtBefore("ChannelMessageUpdate", beforeChannelMessageUpdate);
    initializer.registerRtBefore("ChannelMessageRemove", beforeChannelMessageRemove);
    initializer.registerRtBefore("MatchCreate", beforeMatchCreate);
    initializer.registerRtBefore("MatchDataSend", beforeMatchDataSend);
    initializer.registerRtBefore("MatchJoin", beforeMatchJoin);
    initializer.registerRtBefore("MatchLeave", beforeMatchLeave);
    initializer.registerRtBefore("MatchmakerAdd", beforeMatchmakerAdd);
    initializer.registerRtBefore("MatchmakerRemove", beforeMatchmakerRemove);
    initializer.registerRtBefore("StatusFollow", beforeStatusFollow);
    initializer.registerRtBefore("StatusUnfollow", beforeStatusUnfollow);
    initializer.registerRtBefore("StatusUpdate", beforeStatusUpdate);
    initializer.registerRtBefore("Ping", beforePing);
    initializer.registerRtBefore("Pong", beforePong);
    initializer.registerRpc("healthcheck", rpcHealthCheck);
    nk.httpRequest("http://redpanda:8082/topics/nakama", "post", {
        "Content-Type": "application/vnd.kafka.json.v2+json",
    }, JSON.stringify({
        records: [{ key: "nakama", value: { message: "hi", user: "nakama" } }],
    }));
};
var rpcHealthCheck = function (_ctx, logger, _nk, _initializer) {
    logger.info("rpcHealthCheck called");
    return JSON.stringify({ success: true });
};
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
var afterChannelMessageRemove = function (ctx, logger, nk, output, input) {
    if (typeof input
        .channelMessageRemove !== "undefined" &&
        output !== null &&
        typeof output
            .channelMessageRemove !== "undefined") {
    }
};
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
var beforeChannelMessageRemove = function (ctx, logger, nk, envlope) {
    if (typeof envlope
        .channelMessageRemove !== "undefined") {
    }
};
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
var beforePing = function (ctx, logger, nk, envlope) {
    if (typeof envlope.ping !== "undefined") {
    }
};
var beforePong = function (ctx, logger, nk, envlope) {
    if (typeof envlope.pong !== "undefined") {
    }
};
