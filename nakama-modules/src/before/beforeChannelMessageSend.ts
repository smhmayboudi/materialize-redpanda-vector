const beforeChannelMessageSend: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageSend)
      .channelMessageSend !== "undefined"
  ) {
  }
};
