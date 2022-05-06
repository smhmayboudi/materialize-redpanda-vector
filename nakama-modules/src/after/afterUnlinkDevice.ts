const afterUnlinkDevice: nkruntime.AfterHookFunction<void, nkruntime.AccountDevice> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkDevice", data, request });
};

