const afterAuthenticateCustom: nkruntime.AfterHookFunction<nkruntime.Session, nkruntime.AuthenticateCustomRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAuthenticateCustom", data, request });
};
