const afterBlockFriends: nkruntime.AfterHookFunction<void, nkruntime.BlockFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterBlockFriends", data, request });
};
