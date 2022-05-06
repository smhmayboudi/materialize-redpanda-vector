const afterKickGroupUsers: nkruntime.AfterHookFunction<void, nkruntime.KickGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterKickGroupUsers", data, request });
};

