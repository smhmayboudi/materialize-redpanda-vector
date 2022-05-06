const beforeUnlinkApple: nkruntime.BeforeHookFunction<nkruntime.AccountApple> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkApple", data });
  return data;
};
