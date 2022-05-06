const beforeListStorageObjects: nkruntime.BeforeHookFunction<nkruntime.ListStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListStorageObjects", data });
  return data;
};
