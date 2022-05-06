const afterListChannelMessages: nkruntime.AfterHookFunction<nkruntime.ChannelMessageList, nkruntime.ListChannelMessagesRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterListChannelMessages", data, request });
};

