const afterDeleteNotification: nkruntime.AfterHookFunction<void, nkruntime.DeleteNotificationsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDeleteNotification", data, request });
};
