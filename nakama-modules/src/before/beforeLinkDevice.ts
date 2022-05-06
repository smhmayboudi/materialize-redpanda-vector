const beforeLinkDevice: nkruntime.BeforeHookFunction<nkruntime.AccountDevice> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkDevice", data });
  return data;
};
