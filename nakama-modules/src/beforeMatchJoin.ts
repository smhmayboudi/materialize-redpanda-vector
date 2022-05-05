const beforeMatchJoin: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined"
  ) {
  }
};
