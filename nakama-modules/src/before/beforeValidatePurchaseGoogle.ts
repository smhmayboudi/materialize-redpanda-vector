const beforeValidatePurchaseGoogle: nkruntime.BeforeHookFunction<nkruntime.ValidatePurchaseGoogleRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeValidatePurchaseGoogle", data });
  return data;
};
