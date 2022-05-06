const afterStatusUpdate: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeStatusUpdate).statusUpdate !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusUpdate).statusUpdate !==
    "undefined"
  ) {
  }
};
