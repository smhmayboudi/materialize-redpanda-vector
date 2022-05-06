const afterAuthenticateFacebook: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateFacebookRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateFacebook", data, request });
};
