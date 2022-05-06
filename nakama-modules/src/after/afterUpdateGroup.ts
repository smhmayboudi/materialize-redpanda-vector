const afterUpdateGroup: nkruntime.AfterHookFunction<void, nkruntime.UpdateGroupRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUpdateGroup", data, request });
};

