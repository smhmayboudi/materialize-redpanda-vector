const beforeListNotifications: nkruntime.BeforeHookFunction<nkruntime.ListNotificationsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListNotifications", data });
  return data;
};
