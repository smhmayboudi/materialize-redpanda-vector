const afterChannelMessageRemove: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  if (
    typeof (input as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined"
  ) {
  }
};
