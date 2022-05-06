const afterWriteLeaderboardRecord: nkruntime.AfterHookFunction<nkruntime.LeaderboardRecord, nkruntime.WriteLeaderboardRecordRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterWriteLeaderboardRecord", data, request });
};

