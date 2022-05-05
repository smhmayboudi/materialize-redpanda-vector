const beforePong: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (typeof (envlope as nkruntime.EnvelopePong).pong !== "undefined") {
  }
};
