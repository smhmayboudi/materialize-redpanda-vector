const beforeListTournamentRecords: nkruntime.BeforeHookFunction<nkruntime.ListTournamentRecordsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListTournamentRecords", data });
  return data;
};
