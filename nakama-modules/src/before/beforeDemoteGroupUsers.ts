const beforeDemoteGroupUsers: nkruntime.BeforeHookFunction<nkruntime.DemoteGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDemoteGroupUsers", data });
  return data;
};
