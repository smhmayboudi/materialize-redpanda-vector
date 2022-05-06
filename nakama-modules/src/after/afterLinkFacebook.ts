const afterLinkFacebook: nkruntime.AfterHookFunction<void, nkruntime.LinkFacebookRequest> = (
  ctx,
  logger,
  nk,
  data,
  request
) => {
  redpanda(ctx, logger, nk, { name: "afterLinkFacebook", data, request });
};

