const beforeGetUsers: nkruntime.BeforeHookFunction<nkruntime.GetUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeGetUsers", data });
  return data;
};
