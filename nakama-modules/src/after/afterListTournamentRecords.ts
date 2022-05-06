const afterListTournamentRecords: nkruntime.AfterHookFunction<nkruntime.TournamentRecordList, nkruntime.ListTournamentRecordsRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListTournamentRecords", data, request });
};

