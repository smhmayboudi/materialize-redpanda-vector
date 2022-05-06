const beforeMatchCreate: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchCreate", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchCreateMessage).matchCreate !==
    "undefined"
  ) {
  }
  return envlope;
};
