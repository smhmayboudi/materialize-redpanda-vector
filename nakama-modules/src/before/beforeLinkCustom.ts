const beforeLinkCustom: nkruntime.BeforeHookFunction<nkruntime.AccountCustom> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkCustom", data });
  return data;
};
