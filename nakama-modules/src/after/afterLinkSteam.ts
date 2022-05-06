const afterLinkSteam: nkruntime.AfterHookFunction<void, nkruntime.LinkSteamRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkSteam", data, request });
};

