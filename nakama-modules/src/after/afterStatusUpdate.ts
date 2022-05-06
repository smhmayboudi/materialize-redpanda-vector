const afterStatusUpdate: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterStatusUpdate", output, input });
  if (
    typeof (input as nkruntime.EnvelopeStatusUpdate).statusUpdate !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusUpdate).statusUpdate !==
    "undefined"
  ) {
  }
  return output;
};
