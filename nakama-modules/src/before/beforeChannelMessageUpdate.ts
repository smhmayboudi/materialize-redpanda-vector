const beforeChannelMessageUpdate: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined"
  ) {
  }
};
