const afterListGroupUsers: nkruntime.AfterHookFunction<nkruntime.GroupUserList, nkruntime.ListGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListGroupUsers", data, request });
};

