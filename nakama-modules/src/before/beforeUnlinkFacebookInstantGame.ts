const beforeUnlinkFacebookInstantGame: nkruntime.BeforeHookFunction<nkruntime.AccountFacebook> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkFacebookInstantGame", data });
  return data;
};
