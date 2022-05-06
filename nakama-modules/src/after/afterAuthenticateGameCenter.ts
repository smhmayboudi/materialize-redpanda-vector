const afterAuthenticateGameCenter: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateGameCenterRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateGameCenter", data, request });
};
