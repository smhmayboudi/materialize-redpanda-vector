const beforeAuthenticateFacebookInstantGame: nkruntime.BeforeHookFunction<nkruntime.AuthenticateFacebookInstantGameRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateFacebookInstantGame", data });
  return data;
};
