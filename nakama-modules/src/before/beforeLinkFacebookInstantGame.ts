const beforeLinkFacebookInstantGame: nkruntime.BeforeHookFunction<nkruntime.AccountFacebookInstantGame> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkFacebookInstantGame", data });
  return data;
};
