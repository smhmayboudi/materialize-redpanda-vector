const beforeImportSteamFriends: nkruntime.BeforeHookFunction<nkruntime.ImportSteamFriendsRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeImportSteamFriends", data });
  return data;
};
