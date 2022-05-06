const afterListLeaderboardRecords: nkruntime.AfterHookFunction<nkruntime.LeaderboardRecordList, nkruntime.ListLeaderboardRecordsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListLeaderboardRecords", data, request });
};

