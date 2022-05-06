const registerLeaderboardReset: nkruntime.LeaderboardResetFunction = (
    ctx,
    logger,
    nk,
    leaderboard,
    reset) => {
    redpanda(ctx, logger, nk, { name: "registerLeaderboardReset", leaderboard, reset });
};
