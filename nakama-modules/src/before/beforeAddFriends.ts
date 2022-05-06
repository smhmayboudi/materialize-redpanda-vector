const beforeAddFriends: nkruntime.BeforeHookFunction<nkruntime.AddFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeAddFriends", data });
  return data;
};
