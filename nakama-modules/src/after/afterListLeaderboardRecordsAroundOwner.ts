const afterListLeaderboardRecordsAroundOwner: nkruntime.AfterHookFunction<nkruntime.LeaderboardRecordList, nkruntime.ListLeaderboardRecordsAroundOwnerRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListLeaderboardRecordsAroundOwner", data, request });
};

