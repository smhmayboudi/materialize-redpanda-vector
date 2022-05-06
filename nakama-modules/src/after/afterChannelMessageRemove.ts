const afterChannelMessageRemove: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  redpanda(ctx, logger, nk, { name: "afterChannelMessageRemove", output, input });
  if (
    typeof (input as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelMessageRemove)
      .channelMessageRemove !== "undefined"
  ) {
  }
  return output;
};
