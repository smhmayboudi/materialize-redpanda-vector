const afterWriteTournamentRecord: nkruntime.AfterHookFunction<nkruntime.LeaderboardRecord, nkruntime.WriteTournamentRecordRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterWriteTournamentRecord", data, request });
};

