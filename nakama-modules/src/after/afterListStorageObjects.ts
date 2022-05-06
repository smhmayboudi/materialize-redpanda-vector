const afterListStorageObjects: nkruntime.AfterHookFunction<nkruntime.StorageObjectList, nkruntime.ListStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListStorageObjects", data, request });
};

