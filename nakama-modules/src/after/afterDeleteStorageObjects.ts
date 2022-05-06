const afterDeleteStorageObjects: nkruntime.AfterHookFunction<void, nkruntime.DeleteStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDeleteStorageObjects", data, request });
};

