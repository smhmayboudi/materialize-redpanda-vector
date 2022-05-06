const beforeMatchJoin: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchJoin", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined"
  ) {
  }
  return envlope;
};
