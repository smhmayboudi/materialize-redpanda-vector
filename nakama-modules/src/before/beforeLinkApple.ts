const beforeLinkApple: nkruntime.BeforeHookFunction<nkruntime.AccountApple> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkApple", data });
  return data;
};
