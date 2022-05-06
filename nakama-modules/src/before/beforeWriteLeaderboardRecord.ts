const beforeWriteLeaderboardRecord: nkruntime.BeforeHookFunction<nkruntime.WriteLeaderboardRecordRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeWriteLeaderboardRecord", data });
  return data;
};
