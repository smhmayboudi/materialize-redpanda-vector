const tournamentEnd: nkruntime.TournamentEndFunction = (
    ctx,
    logger,
    nk,
    tournament,
    end,
    reset) => {
    redpanda(ctx, logger, nk, { name: "tournamentEnd", tournament, end, reset });
};
