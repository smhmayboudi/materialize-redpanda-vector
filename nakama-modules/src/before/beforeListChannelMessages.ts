const beforeListChannelMessages: nkruntime.BeforeHookFunction<nkruntime.ListChannelMessagesRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeListChannelMessages", data });
  return data;
};
