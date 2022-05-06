const afterMatchCreate: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchCreateMessage).matchCreate !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchCreateMessage).matchCreate !==
    "undefined"
  ) {
  }
};
