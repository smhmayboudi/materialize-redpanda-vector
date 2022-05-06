const afterDeleteFriends: nkruntime.AfterHookFunction<void, nkruntime.DeleteFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterDeleteFriends", data, request });
};
