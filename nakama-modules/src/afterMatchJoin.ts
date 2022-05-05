const afterMatchJoin: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined"
  ) {
  }
};
