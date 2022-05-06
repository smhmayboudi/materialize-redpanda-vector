const afterCreateGroup: nkruntime.AfterHookFunction<nkruntime.Group, nkruntime.CreateGroupRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterCreateGroup", data, request });
};
