const afterLeaveGroup: nkruntime.AfterHookFunction<void, nkruntime.LeaveGroupRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLeaveGroup", data, request });
};

