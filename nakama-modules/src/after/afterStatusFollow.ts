const afterStatusFollow: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined"
  ) {
  }
};
