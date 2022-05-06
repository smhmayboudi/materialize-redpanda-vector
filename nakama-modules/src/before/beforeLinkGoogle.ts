const beforeLinkGoogle: nkruntime.BeforeHookFunction<nkruntime.AccountGoogle> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkGoogle", data });
  return data;
};
