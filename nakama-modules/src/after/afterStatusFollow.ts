const afterStatusFollow: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterStatusFollow", output, input });
  if (
    typeof (input as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopeStatusFollow).statusFollow !==
    "undefined"
  ) {
  }
  return output;
};
