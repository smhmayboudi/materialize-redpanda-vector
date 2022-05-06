const beforeAuthenticateGameCenter: nkruntime.BeforeHookFunction<nkruntime.AuthenticateGameCenterRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateGameCenter", data });
  return data;
};
