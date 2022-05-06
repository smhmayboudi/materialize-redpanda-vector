const afterUnlinkEmail: nkruntime.AfterHookFunction<void, nkruntime.AccountEmail> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUnlinkEmail", data, request });
};

