const afterValidatePurchaseApple: nkruntime.AfterHookFunction<nkruntime.ValidatePurchaseResponse, nkruntime.ValidatePurchaseAppleRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterValidatePurchaseApple", data, request });
};

