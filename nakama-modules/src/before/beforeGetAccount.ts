const beforeGetAccount: nkruntime.BeforeHookFunction<void> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeGetAccount", data });
  return data;
};
