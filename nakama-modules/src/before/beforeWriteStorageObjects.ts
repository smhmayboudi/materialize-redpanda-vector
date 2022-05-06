const beforeWriteStorageObjects: nkruntime.BeforeHookFunction<nkruntime.WriteStorageObjectsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeWriteStorageObjects", data });
  return data;
};
