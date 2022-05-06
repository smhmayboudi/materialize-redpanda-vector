const beforeKickGroupUsers: nkruntime.BeforeHookFunction<nkruntime.KickGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeKickGroupUsers", data });
  return data;
};
