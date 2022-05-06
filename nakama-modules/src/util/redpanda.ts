
const redpanda = function (
    ctx: nkruntime.Context,
    logger: nkruntime.Logger,
    nk: nkruntime.Nakama,
    payload: { [key: string]: any }
) {
    nk.httpRequest(
        "http://redpanda:8082/topics/nakama",
        "post",
        {
            "Content-Type": "application/vnd.kafka.json.v2+json"
        },
        JSON.stringify({
            "records": [{
                "key": { "node": ctx.node },
                "value": { ctx, payload }
            }]
        })
    );
};
