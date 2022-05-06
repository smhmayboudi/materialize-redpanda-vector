const afterDeleteGroup: nkruntime.AfterHookFunction<void, nkruntime.DeleteGroupRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDeleteGroup", data, request });
};
