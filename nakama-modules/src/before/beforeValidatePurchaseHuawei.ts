const beforeValidatePurchaseHuawei: nkruntime.BeforeHookFunction<nkruntime.ValidatePurchaseHuaweiRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeValidatePurchaseHuawei", data });
  return data;
};
