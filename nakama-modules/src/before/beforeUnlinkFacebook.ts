const beforeUnlinkFacebook: nkruntime.BeforeHookFunction<nkruntime.AccountFacebookInstantGame> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUnlinkFacebook", data });
  return data;
};
