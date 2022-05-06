const beforeListUserGroups: nkruntime.BeforeHookFunction<nkruntime.ListUserGroupsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListUserGroups", data });
  return data;
};
