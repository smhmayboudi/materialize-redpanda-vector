const afterPing: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterPing", output, input });
  if (
    typeof (input as nkruntime.EnvelopePing).ping !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopePing).ping !== "undefined"
  ) {
  }
  return output;
};
