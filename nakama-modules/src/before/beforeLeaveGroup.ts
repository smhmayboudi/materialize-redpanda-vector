const beforeLeaveGroup: nkruntime.BeforeHookFunction<nkruntime.LeaveGroupRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLeaveGroup", data });
  return data;
};
