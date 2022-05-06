const afterChannelLeave: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterChannelLeave", output, input });
  if (
    typeof (input as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined"
  ) {
  }
  return output;
};
