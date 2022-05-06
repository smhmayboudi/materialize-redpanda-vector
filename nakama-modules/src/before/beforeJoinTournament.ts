const beforeJoinTournament: nkruntime.BeforeHookFunction<nkruntime.JoinTournamentRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeJoinTournament", data });
  return data;
};
