const afterUpdateAccount: nkruntime.AfterHookFunction<void, nkruntime.UserUpdateAccount> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterUpdateAccount", data, request });
};

