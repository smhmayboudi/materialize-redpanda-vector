const afterDeleteLeaderboardRecord: nkruntime.AfterHookFunction<void, nkruntime.DeleteLeaderboardRecordRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDeleteLeaderboardRecord", data, request });
};
