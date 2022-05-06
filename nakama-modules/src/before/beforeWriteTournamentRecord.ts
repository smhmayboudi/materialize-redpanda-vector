const beforeWriteTournamentRecord: nkruntime.BeforeHookFunction<nkruntime.WriteTournamentRecordRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeWriteTournamentRecord", data });
  return data;
};
