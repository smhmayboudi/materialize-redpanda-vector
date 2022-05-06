const afterStatusUnfollow: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterStatusUnfollow", output, input });
  if (
    typeof (input as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined"
  ) {
  }
  return output;
};
