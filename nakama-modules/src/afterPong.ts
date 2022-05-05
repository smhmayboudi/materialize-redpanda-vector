const afterPong: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopePong).pong !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopePong).pong !== "undefined"
  ) {
  }
};
