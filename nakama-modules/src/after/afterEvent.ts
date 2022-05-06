const afterEvent: nkruntime.AfterHookFunction<void, nkruntime.Event> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterEvent", data, request });
};

