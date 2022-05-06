const beforeMatchmakerAdd: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined"
  ) {
  }
};
