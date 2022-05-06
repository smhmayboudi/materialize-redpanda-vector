const beforeMatchmakerAdd: nkruntime.RtBeforeHookFunction<
  nkruntime.Envelope
> = (ctx, logger, nk, envlope) => {
  // redpanda(ctx, logger, nk, { name: "beforeMatchmakerAdd", envlope });
  if (
    typeof (envlope as nkruntime.EnvelopeMatchmakerAdd).matchmakerAdd !==
    "undefined"
  ) {
  }
  return envlope;
};
