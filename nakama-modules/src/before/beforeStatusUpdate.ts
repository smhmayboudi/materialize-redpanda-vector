const beforeStatusUpdate: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeStatusUpdate", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeStatusUpdate).statusUpdate !==
    "undefined"
  ) {
  }
  return envlope;
};
