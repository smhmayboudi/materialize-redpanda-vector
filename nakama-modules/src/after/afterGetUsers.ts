const afterGetUsers: nkruntime.AfterHookFunction<nkruntime.Users, nkruntime.GetUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterGetUsers", data, request });
};

