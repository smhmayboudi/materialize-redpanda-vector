const beforeAuthenticateSteam: nkruntime.BeforeHookFunction<nkruntime.AuthenticateSteamRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateSteam", data });
  return data;
};
