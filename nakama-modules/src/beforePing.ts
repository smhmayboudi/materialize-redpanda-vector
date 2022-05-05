const beforePing: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (typeof (envlope as nkruntime.EnvelopePing).ping !== "undefined") {
  }
};
