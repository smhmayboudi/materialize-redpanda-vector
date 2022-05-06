const afterAddGroupUsers: nkruntime.AfterHookFunction<void, nkruntime.AddGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAddGroupUsers", data, request });
};
