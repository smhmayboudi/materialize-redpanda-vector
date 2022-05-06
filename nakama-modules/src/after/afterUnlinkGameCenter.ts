const afterUnlinkGameCenter: nkruntime.AfterHookFunction<void, nkruntime.AccountGameCenter> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkGameCenter", data, request });
};

