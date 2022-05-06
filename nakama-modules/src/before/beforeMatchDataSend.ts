const beforeMatchDataSend: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchDataSend).matchDataSend !==
    "undefined"
  ) {
  }
};
