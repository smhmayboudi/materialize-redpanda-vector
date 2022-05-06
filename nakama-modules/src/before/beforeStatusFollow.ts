const beforeStatusFollow: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined"
  ) {
  }
};
