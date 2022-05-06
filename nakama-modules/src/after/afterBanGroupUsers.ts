const afterBanGroupUsers: nkruntime.AfterHookFunction<void, nkruntime.BanGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterBanGroupUsers", data, request });
};
