const beforeMatchCreate: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchCreateMessage).matchCreate !==
    "undefined"
  ) {
  }
};
