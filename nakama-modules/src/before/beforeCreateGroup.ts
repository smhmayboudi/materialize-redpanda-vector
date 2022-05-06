const beforeCreateGroup: nkruntime.BeforeHookFunction<nkruntime.CreateGroupRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeCreateGroup", data });
  return data;
};
