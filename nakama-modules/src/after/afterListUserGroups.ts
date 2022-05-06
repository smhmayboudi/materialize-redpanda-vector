const afterListUserGroups: nkruntime.AfterHookFunction<nkruntime.UserGroupList, nkruntime.ListUserGroupsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListUserGroups", data, request });
};

