const afterChannelMessageSend: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  redpanda(ctx, logger, nk, { name: "afterChannelMessageSend", output, input });
  if (
    typeof (input as nkruntime.EnvelopeChannelMessageSend)
      .channelMessageSend !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelMessageSend)
      .channelMessageSend !== "undefined"
  ) {
  }
  return output;
};
