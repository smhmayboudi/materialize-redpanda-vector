const beforeStatusUnfollow: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  if (
    typeof (envlope as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined"
  ) {
  }
};
