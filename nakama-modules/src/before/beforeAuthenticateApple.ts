const beforeAuthenticateApple: nkruntime.BeforeHookFunction<nkruntime.AuthenticateAppleRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateApple", data });
  return data;
};
