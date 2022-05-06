const afterDemoteGroupUsers: nkruntime.AfterHookFunction<void, nkruntime.DemoteGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDemoteGroupUsers", data, request });
};

