const afterUnlinkFacebookInstantGame: nkruntime.AfterHookFunction<void, nkruntime.AccountFacebookInstantGame> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkFacebookInstantGame", data, request });
};

