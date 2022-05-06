const beforeChannelMessageRemove: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined"
  ) {
  }
};
