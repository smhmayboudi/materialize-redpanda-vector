const beforeUpdateAccount: nkruntime.BeforeHookFunction<nkruntime.UserUpdateAccount> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeUpdateAccount", data });
  return data;
};
