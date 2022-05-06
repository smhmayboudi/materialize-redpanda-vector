const beforeEvent: nkruntime.BeforeHookFunction<nkruntime.Event> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeEvent", data });
  return data;
};
