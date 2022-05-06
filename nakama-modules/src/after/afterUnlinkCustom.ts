const afterUnlinkCustom: nkruntime.AfterHookFunction<void, nkruntime.AccountCustom> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkCustom", data, request });
};

