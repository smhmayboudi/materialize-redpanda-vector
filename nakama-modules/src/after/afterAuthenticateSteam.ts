const afterAuthenticateSteam: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateSteamRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateSteam", data, request });
};
