const matchmakerMatched: nkruntime.MatchmakerMatchedFunction = (
    ctx,
    logger,
    nk,
    matches) => {
    redpanda(ctx, logger, nk, { name: "matchmakerMatched", matches });
};
