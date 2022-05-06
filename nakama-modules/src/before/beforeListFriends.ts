const beforeListFriends: nkruntime.BeforeHookFunction<nkruntime.ListFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListFriends", data });
  return data;
};
