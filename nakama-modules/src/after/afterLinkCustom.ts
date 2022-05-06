const afterLinkCustom: nkruntime.AfterHookFunction<void, nkruntime.AccountCustom> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkCustom", data, request });
};

