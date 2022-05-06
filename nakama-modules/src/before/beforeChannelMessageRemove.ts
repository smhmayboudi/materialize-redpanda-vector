const beforeChannelMessageRemove: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeChannelMessageRemove", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined"
  ) {
  }
  return envlope;
};
