const beforeChannelJoin: nkruntime.RtBeforeHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  envlope
) => {
  // redpanda(ctx, logger, nk, { name: "beforeChannelJoin", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeChannelJoin).channelJoin !==
    "undefined"
  ) {
  }
  return envlope;
};
