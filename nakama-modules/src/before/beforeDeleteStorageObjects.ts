const beforeDeleteStorageObjects: nkruntime.BeforeHookFunction<nkruntime.DeleteStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDeleteStorageObjects", data });
  return data;
};
