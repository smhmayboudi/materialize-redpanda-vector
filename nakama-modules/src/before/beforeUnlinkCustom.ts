const beforeUnlinkCustom: nkruntime.BeforeHookFunction<nkruntime.AccountCustom> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkCustom", data });
  return data;
};
