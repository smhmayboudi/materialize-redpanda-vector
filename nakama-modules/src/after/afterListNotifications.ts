const afterListNotifications: nkruntime.AfterHookFunction<nkruntime.NotificationList, nkruntime.ListNotificationsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListNotifications", data, request });
};

