const afterLinkGoogle: nkruntime.AfterHookFunction<void, nkruntime.AccountGoogle> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkGoogle", data, request });
};

