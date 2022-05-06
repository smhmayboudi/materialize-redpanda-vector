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
            "content-type": "application/vnd.kafka.json.v2+json"
        },
        JSON.stringify({
            "records": [{
                "key": {
                    "node": ctx.node
                },
                "value": {
                    context: {
                        client_ip: ctx.clientIp,
                        client_port: ctx.clientPort,
                        env: ctx.env,
                        execution_mode: ctx.executionMode,
                        headers: ctx.headers,
                        lang: ctx.lang,
                        match_id: ctx.matchId,
                        match_label: ctx.matchLabel,
                        match_node: ctx.matchNode,
                        match_tick_rate: ctx.matchTickRate,
                        node: ctx.node,
                        query_params: ctx.queryParams,
                        session_id: ctx.sessionId,
                        user_id: ctx.userId,
                        user_session_exp: ctx.userSessionExp,
                        username: ctx.username,
                        vars: ctx.vars
                    },
                    payload
                }
            }]
        })
    );
};
