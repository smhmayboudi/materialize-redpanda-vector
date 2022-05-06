const beforeBanGroupUsers: nkruntime.BeforeHookFunction<nkruntime.BanGroupUsersRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeBanGroupUsers", data });
  return data;
};
