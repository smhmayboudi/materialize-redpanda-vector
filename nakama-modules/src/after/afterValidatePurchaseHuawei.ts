const afterValidatePurchaseHuawei: nkruntime.AfterHookFunction<nkruntime.ValidatePurchaseResponse, nkruntime.ValidatePurchaseHuaweiRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterValidatePurchaseHuawei", data, request });
};

