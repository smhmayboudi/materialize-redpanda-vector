const afterListGroups: nkruntime.AfterHookFunction<nkruntime.GroupList, nkruntime.ListGroupsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListGroups", data, request });
};

