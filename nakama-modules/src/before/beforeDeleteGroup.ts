const beforeDeleteGroup: nkruntime.BeforeHookFunction<nkruntime.DeleteGroupRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDeleteGroup", data });
  return data;
};
