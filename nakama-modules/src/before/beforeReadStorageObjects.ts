const beforeReadStorageObjects: nkruntime.BeforeHookFunction<nkruntime.ReadStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeReadStorageObjects", data });
  return data;
};
