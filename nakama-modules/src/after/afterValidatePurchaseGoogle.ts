const afterValidatePurchaseGoogle: nkruntime.AfterHookFunction<nkruntime.ValidatePurchaseResponse, nkruntime.ValidatePurchaseGoogleRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterValidatePurchaseGoogle", data, request });
};

