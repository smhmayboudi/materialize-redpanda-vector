const beforeUnlinkEmail: nkruntime.BeforeHookFunction<nkruntime.AccountEmail> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkEmail", data });
  return data;
};
