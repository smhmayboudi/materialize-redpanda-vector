const beforeChannelJoin: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeChannelJoin).channelJoin !==
    "undefined"
  ) {
  }
};
