const afterMatchJoin: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterMatchJoin", output, input });
  if (
    typeof (input as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeMatchJoin).matchJoin !== "undefined"
  ) {
  }
  return output;
};
