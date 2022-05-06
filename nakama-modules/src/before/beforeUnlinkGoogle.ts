const beforeUnlinkGoogle: nkruntime.BeforeHookFunction<nkruntime.AccountGoogle> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkGoogle", data });
  return data;
};
