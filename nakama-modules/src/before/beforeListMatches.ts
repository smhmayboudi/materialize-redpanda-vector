const beforeListMatches: nkruntime.BeforeHookFunction<nkruntime.ListMatchesRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListMatches", data });
  return data;
};
