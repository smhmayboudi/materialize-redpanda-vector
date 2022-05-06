const beforeAddGroupUsers: nkruntime.BeforeHookFunction<nkruntime.AddGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAddGroupUsers", data });
  return data;
};
