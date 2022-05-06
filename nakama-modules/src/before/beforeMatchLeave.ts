const beforeMatchLeave: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchLeave", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchLeave).matchLeave !== "undefined"
  ) {
  }
  return envlope;
};
