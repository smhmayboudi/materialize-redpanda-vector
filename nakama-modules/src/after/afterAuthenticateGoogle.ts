const afterAuthenticateGoogle: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateGoogleRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateGoogle", data, request });
};
