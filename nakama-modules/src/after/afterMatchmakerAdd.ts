const afterMatchmakerAdd: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterMatchmakerAdd", output, input });
  if (
    typeof (input as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined"
  ) {
  }
  return output;
};
