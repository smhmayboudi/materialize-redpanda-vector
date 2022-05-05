const afterMatchmakerRemove: nkruntime.RtAfterHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, output, input) => {
  if (
    typeof (input as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined"
  ) {
  }
};
