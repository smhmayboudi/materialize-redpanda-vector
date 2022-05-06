const beforeStatusFollow: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeStatusFollow", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined"
  ) {
  }
  return envlope;
};
