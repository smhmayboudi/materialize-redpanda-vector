const beforeAuthenticateFacebook: nkruntime.BeforeHookFunction<nkruntime.AuthenticateFacebookRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateFacebook", data });
  return data;
};
