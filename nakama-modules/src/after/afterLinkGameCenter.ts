const afterLinkGameCenter: nkruntime.AfterHookFunction<void, nkruntime.AccountGameCenter> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkGameCenter", data, request });
};

