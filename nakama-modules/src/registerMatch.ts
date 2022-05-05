
// registerMatch < State = MatchState > (name: string, functions: MatchHandler<State>): void;
const matchInit: nkruntime.MatchInitFunction = (ctx, logger, nk, params) => {
    return { state: {}, tickRate: 0, label: "" };
};
const matchJoinAttempt: nkruntime.MatchJoinAttemptFunction = (ctx, logger, nk, dispatcher, tick, state, presence, metadata) => {
    return { state: {}, accept: false, rejectMessage: "" };
};
const matchJoin: nkruntime.MatchJoinFunction = (ctx, logger, nk, dispatcher, tick, state, presences) => {
    return { state: {} };
};
const matchLeave: nkruntime.MatchLeaveFunction = (ctx, logger, nk, dispatcher, tick, state, presences) => {
    return { state: {} };
};
const matchLoop: nkruntime.MatchLoopFunction = (ctx, logger, nk, dispatcher, tick, state, messages) => {
    return { state: {} };
};
const matchTerminate: nkruntime.MatchTerminateFunction = (ctx, logger, nk, dispatcher, tick, state, graceSeconds) => {
    return { state: {} };
};
const matchSignal: nkruntime.MatchSignalFunction = (ctx, logger, nk, dispatcher, tick, state, data) => {
    return { state: {}, date: String(new Date()) };
};

const match: nkruntime.MatchHandler = {
    matchInit,
    matchJoinAttempt,
    matchJoin,
    matchLeave,
    matchLoop,
    matchTerminate,
    matchSignal
}
