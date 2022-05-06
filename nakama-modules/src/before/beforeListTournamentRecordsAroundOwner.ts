const beforeListTournamentRecordsAroundOwner: nkruntime.BeforeHookFunction<nkruntime.ListTournamentRecordsAroundOwnerRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListTournamentRecordsAroundOwner", data });
  return data;
};
