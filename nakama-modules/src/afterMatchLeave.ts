const afterMatchLeave: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchLeave).matchLeave !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchLeave).matchLeave !== "undefined"
  ) {
  }
};
