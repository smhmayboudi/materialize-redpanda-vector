const beforeAuthenticateGoogle: nkruntime.BeforeHookFunction<nkruntime.AuthenticateGoogleRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateGoogle", data });
  return data;
};
