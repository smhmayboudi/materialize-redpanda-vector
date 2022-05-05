const afterPing: nkruntime.RtAfterHookFunction<nkruntime.Envelope> = (
  ctx,
  logger,
  nk,
  output,
  input
) => {
  if (
    typeof (input as nkruntime.EnvelopePing).ping !== "undefined" &&
    output !== null &&
    typeof (output as nkruntime.EnvelopePing).ping !== "undefined"
  ) {
  }
};
