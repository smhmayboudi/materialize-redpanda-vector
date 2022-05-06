const afterLinkApple: nkruntime.AfterHookFunction<void, nkruntime.AccountApple> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkApple", data, request });
};

