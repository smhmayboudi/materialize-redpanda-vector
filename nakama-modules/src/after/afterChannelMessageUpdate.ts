const afterChannelMessageUpdate: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  redpanda(ctx, logger, nk, { name: "afterChannelMessageUpdate", output, input });
  if (
    typeof (input as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelMessageUpdate)
      .channelMessageUpdate !== "undefined"
  ) {
  }
  return output;
};
