const beforeChannelMessageUpdate: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeChannelMessageUpdate", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined"
  ) {
  }
  return envlope;
};
