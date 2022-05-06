const afterPromoteGroupUsers: nkruntime.AfterHookFunction<void, nkruntime.PromoteGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterPromoteGroupUsers", data, request });
};

