const afterListTournamentRecordsAroundOwner: nkruntime.AfterHookFunction<nkruntime.LeaderboardRecordList, nkruntime.ListTournamentRecordsAroundOwnerRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListTournamentRecordsAroundOwner", data, request });
};

