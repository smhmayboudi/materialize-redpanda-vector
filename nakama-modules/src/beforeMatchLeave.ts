const beforeMatchLeave: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchLeave).matchLeave !== "undefined"
  ) {
  }
};
