const afterChannelMessageUpdate: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  if (
    typeof (input as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined"
  ) {
  }
};
