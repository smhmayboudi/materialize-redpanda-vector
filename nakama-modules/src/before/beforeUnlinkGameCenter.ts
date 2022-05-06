const beforeUnlinkGameCenter: nkruntime.BeforeHookFunction<nkruntime.AccountGameCenter> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkGameCenter", data });
  return data;
};
