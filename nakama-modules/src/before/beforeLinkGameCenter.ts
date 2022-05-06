const beforeLinkGameCenter: nkruntime.BeforeHookFunction<nkruntime.AccountGameCenter> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkGameCenter", data });
  return data;
};
