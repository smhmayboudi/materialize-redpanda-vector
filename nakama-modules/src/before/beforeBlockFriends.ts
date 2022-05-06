const beforeBlockFriends: nkruntime.BeforeHookFunction<nkruntime.BlockFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeBlockFriends", data });
  return data;
};
