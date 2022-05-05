const rpcHealthCheck: nkruntime.RpcFunction = function (
    _ctx,
    logger,
    _nk,
    _initializer
) {
    logger.info("rpcHealthCheck called");
    return JSON.stringify({ success: true });
};
