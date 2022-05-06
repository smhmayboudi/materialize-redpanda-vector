const afterImportFacebookFriends: nkruntime.AfterHookFunction<void, nkruntime.ImportFacebookFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterImportFacebookFriends", data, request });
};

