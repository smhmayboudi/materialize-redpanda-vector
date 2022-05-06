const beforeDeleteNotification: nkruntime.BeforeHookFunction<nkruntime.DeleteNotificationsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDeleteNotification", data });
  return data;
};
