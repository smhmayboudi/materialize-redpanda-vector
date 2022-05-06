const beforeListLeaderboardRecords: nkruntime.BeforeHookFunction<nkruntime.ListLeaderboardRecordsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListLeaderboardRecords", data });
  return data;
};
