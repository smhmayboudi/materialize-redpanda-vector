const afterListMatches: nkruntime.AfterHookFunction<nkruntime.MatchList, nkruntime.ListMatchesRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListMatches", data, request });
};

