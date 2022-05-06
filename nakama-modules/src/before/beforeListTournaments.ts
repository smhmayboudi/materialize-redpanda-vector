const beforeListTournaments: nkruntime.BeforeHookFunction<nkruntime.ListTournamentsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListTournaments", data });
  return data;
};
