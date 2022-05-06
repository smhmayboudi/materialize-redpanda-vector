const beforeChannelMessageSend: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeChannelMessageSend", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageSend)
      .channelMessageSend !== "undefined"
  ) {
  }
  return envlope;
};
