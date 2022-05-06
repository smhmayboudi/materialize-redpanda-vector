const afterUnlinkApple: nkruntime.AfterHookFunction<void, nkruntime.AccountApple> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkApple", data, request });
};

