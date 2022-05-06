const beforeLinkEmail: nkruntime.BeforeHookFunction<nkruntime.AccountEmail> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkEmail", data });
  return data;
};
