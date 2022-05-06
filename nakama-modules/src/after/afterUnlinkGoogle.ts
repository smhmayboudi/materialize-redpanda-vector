const afterUnlinkGoogle: nkruntime.AfterHookFunction<void, nkruntime.AccountGoogle> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkGoogle", data, request });
};

