const afterGetAccount: nkruntime.AfterHookFunction<nkruntime.Account, void> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterGetAccount", data, request });
};

