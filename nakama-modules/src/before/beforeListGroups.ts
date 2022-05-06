const beforeListGroups: nkruntime.BeforeHookFunction<nkruntime.ListGroupsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListGroups", data });
  return data;
};
