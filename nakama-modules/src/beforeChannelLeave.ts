const beforeChannelLeave: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  if (
    typeof (envlope as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined"
  ) {
  }
};
