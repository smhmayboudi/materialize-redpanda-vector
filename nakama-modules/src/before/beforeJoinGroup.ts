const beforeJoinGroup: nkruntime.BeforeHookFunction<nkruntime.JoinGroupRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeJoinGroup", data });
  return data;
};
