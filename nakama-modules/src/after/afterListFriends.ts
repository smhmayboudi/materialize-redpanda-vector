const afterListFriends: nkruntime.AfterHookFunction<nkruntime.FriendList, nkruntime.ListFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListFriends", data, request });
};

