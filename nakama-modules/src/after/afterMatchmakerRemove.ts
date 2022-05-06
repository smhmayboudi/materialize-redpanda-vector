const afterMatchmakerRemove: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  redpanda(ctx, logger, nk, { name: "afterMatchmakerRemove", output, input });
  if (
    typeof (input as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined"
  ) {
  }
  return output;
};
