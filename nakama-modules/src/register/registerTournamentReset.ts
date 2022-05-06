const tournamentReset: nkruntime.TournamentResetFunction = (
    ctx,
    logger,
    nk,
    tournament,
    end,
    reset) => {
    redpanda(ctx, logger, nk, { name: "tournamentReset", tournament, end, reset });
};
