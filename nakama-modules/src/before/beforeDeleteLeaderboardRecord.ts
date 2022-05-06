const beforeDeleteLeaderboardRecord: nkruntime.BeforeHookFunction<nkruntime.DeleteLeaderboardRecordRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDeleteLeaderboardRecord", data });
  return data;
};
