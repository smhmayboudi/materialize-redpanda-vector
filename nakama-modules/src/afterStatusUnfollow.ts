const afterStatusUnfollow: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined"
  ) {
  }
};
