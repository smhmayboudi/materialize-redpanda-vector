const beforeUnlinkDevice: nkruntime.BeforeHookFunction<nkruntime.AccountDevice> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkDevice", data });
  return data;
};
