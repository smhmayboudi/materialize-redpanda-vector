const afterWriteStorageObjects: nkruntime.AfterHookFunction<nkruntime.StorageObjectAcks, nkruntime.WriteStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterWriteStorageObjects", data, request });
};

