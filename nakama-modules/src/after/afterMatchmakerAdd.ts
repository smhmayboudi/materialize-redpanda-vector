const afterMatchmakerAdd: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined"
  ) {
  }
};
