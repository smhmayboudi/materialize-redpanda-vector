const beforeChannelLeave: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeChannelLeave", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined"
  ) {
  }
  return envlope;
};
