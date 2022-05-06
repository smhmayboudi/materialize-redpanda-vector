const beforeStatusUnfollow: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeStatusUnfollow", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeStatusUnfollow).statusUnfollow !==
    "undefined"
  ) {
  }
  return envlope;
};
