const afterChannelLeave: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeChannelLeave).channelLeave !==
    "undefined"
  ) {
  }
};
