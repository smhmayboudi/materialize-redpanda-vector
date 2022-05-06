const afterAuthenticateApple: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateAppleRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateApple", data, request });
};
