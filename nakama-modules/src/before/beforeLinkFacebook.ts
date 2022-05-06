const beforeLinkFacebook: nkruntime.BeforeHookFunction<nkruntime.LinkFacebookRequest> = (
  ctx,
  logger,
  nk,
  data
) => {
  // redpanda(ctx, logger, nk, { name: "beforeLinkFacebook", data });
  return data;
};
