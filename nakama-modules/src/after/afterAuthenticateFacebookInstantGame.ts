const afterAuthenticateFacebookInstantGame: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateFacebookInstantGameRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateFacebookInstantGame", data, request });
};
