const afterImportSteamFriends: nkruntime.AfterHookFunction<void, nkruntime.ImportSteamFriendsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterImportSteamFriends", data, request });
};

