const beforeListGroupUsers: nkruntime.BeforeHookFunction<nkruntime.ListGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListGroupUsers", data });
  return data;
};
