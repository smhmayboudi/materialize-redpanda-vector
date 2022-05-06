const beforeValidatePurchaseApple: nkruntime.BeforeHookFunction<nkruntime.ValidatePurchaseAppleRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeValidatePurchaseApple", data });
  return data;
};
