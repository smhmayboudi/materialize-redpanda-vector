const afterChannelJoin: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterChannelJoin", output, input });
  if (
    typeof (input as nkruntime.EnvelopeChannelJoin).channelJoin !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelJoin).channelJoin !== "undefined"
  ) {
  }
  return output;
};
