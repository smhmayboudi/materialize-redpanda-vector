const afterAddFriends: nkruntime.AfterHookFunction<void, nkruntime.AddFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterAddFriends", data, request });
};
