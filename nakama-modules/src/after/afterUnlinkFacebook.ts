const afterUnlinkFacebook: nkruntime.AfterHookFunction<void, nkruntime.AccountFacebook> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkFacebook", data, request });
};

