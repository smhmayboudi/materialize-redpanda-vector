const beforeImportFacebookFriends: nkruntime.BeforeHookFunction<nkruntime.ImportFacebookFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeImportFacebookFriends", data });
  return data;
};
