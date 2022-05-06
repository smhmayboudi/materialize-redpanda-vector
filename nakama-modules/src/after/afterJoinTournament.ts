const afterJoinTournament: nkruntime.AfterHookFunction<void, nkruntime.JoinTournamentRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterJoinTournament", data, request });
};

