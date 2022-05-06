const afterJoinGroup: nkruntime.AfterHookFunction<void, nkruntime.JoinGroupRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterJoinGroup", data, request });
};

