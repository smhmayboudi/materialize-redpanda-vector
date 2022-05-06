const beforeLinkSteam: nkruntime.BeforeHookFunction<nkruntime.LinkSteamRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkSteam", data });
  return data;
};
