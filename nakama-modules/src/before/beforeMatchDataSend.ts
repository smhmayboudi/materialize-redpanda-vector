const beforeMatchDataSend: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchDataSend", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchDataSend).matchDataSend !==
    "undefined"
  ) {
  }
  return envlope;
};
