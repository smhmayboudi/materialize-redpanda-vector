const beforePromoteGroupUsers: nkruntime.BeforeHookFunction<nkruntime.PromoteGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforePromoteGroupUsers", data });
  return data;
};
