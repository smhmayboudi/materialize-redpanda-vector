const beforeMatchmakerRemove: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchmakerRemove", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchmakerRemove).matchmakerRemove !==
    "undefined"
  ) {
  }
  return envlope;
};
