const beforeAuthenticateEmail: nkruntime.BeforeHookFunction<nkruntime.AuthenticateEmailRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateEmail", data });
  return data;
};
