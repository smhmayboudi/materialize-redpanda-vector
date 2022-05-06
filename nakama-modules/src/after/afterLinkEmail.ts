const afterLinkEmail: nkruntime.AfterHookFunction<void, nkruntime.AccountEmail> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkEmail", data, request });
};

