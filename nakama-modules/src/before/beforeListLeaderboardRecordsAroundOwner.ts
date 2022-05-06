const beforeListLeaderboardRecordsAroundOwner: nkruntime.BeforeHookFunction<nkruntime.ListLeaderboardRecordsAroundOwnerRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListLeaderboardRecordsAroundOwner", data });
  return data;
};
