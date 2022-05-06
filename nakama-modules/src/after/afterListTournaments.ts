const afterListTournaments: nkruntime.AfterHookFunction<nkruntime.TournamentList, nkruntime.ListTournamentsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListTournaments", data, request });
};

