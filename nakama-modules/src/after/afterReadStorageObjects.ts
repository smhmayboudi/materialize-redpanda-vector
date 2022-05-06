const afterReadStorageObjects: nkruntime.AfterHookFunction<nkruntime.StorageObjects, nkruntime.ReadStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterReadStorageObjects", data, request });
};

