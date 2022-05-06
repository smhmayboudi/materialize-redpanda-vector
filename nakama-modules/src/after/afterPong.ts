const afterPong: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  redpanda(ctx, logger, nk, { name: "afterPong", output, input });
  if (
    typeof (input as nkruntime.EnvelopePong).pong !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopePong).pong !== "undefined"
  ) {
  }
  return output;
};
