const beforeDeleteFriends: nkruntime.BeforeHookFunction<nkruntime.DeleteFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeDeleteFriends", data });
  return data;
};
