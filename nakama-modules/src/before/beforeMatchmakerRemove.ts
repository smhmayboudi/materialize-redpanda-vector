const beforeMatchmakerRemove: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined"
  ) {
  }
};
