const afterUnlinkSteam: nkruntime.AfterHookFunction<void, nkruntime.AccountSteam> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkSteam", data, request });
};

