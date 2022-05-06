const beforeAuthenticateDevice: nkruntime.BeforeHookFunction<nkruntime.AuthenticateDeviceRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAuthenticateDevice", data });
  return data;
};
