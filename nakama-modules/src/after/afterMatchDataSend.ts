const afterMatchDataSend: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchDataSend).matchDataSend !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchDataSend).matchDataSend !==
    "undefined"
  ) {
  }
};
