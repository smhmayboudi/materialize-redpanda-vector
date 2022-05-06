const beforeUpdateGroup: nkruntime.BeforeHookFunction<nkruntime.UpdateGroupRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUpdateGroup", data });
  return data;
};
