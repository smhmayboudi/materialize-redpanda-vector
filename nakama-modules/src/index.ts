
const InitModule: nkruntime.InitModule = function (
  _ctx,
  logger,
  nk,
  initializer
) {
  logger.info("InitModule called");

  initializer.registerAfterAddFriends

  initializer.registerRtAfter("ChannelJoin", afterChannelJoin);
  initializer.registerRtAfter("ChannelLeave", afterChannelLeave);
  initializer.registerRtAfter("ChannelMessageSend", afterChannelMessageSend);
  initializer.registerRtAfter("ChannelMessageUpdate", afterChannelMessageUpdate);
  initializer.registerRtAfter("ChannelMessageRemove", afterChannelMessageRemove);
  initializer.registerRtAfter("MatchCreate", afterMatchCreate);
  initializer.registerRtAfter("MatchDataSend", afterMatchDataSend);
  initializer.registerRtAfter("MatchJoin", afterMatchJoin);
  initializer.registerRtAfter("MatchLeave", afterMatchLeave);
  initializer.registerRtAfter("MatchmakerAdd", afterMatchmakerAdd);
  initializer.registerRtAfter("MatchmakerRemove", afterMatchmakerRemove);
  initializer.registerRtAfter("StatusFollow", afterStatusFollow);
  initializer.registerRtAfter("StatusUnfollow", afterStatusUnfollow);
  initializer.registerRtAfter("StatusUpdate", afterStatusUpdate);
  initializer.registerRtAfter("Ping", afterPing);
  initializer.registerRtAfter("Pong", afterPong);

  initializer.registerRtBefore("ChannelJoin", beforeChannelJoin);
  initializer.registerRtBefore("ChannelLeave", beforeChannelLeave);
  initializer.registerRtBefore("ChannelMessageSend", beforeChannelMessageSend);
  initializer.registerRtBefore("ChannelMessageUpdate", beforeChannelMessageUpdate);
  initializer.registerRtBefore("ChannelMessageRemove", beforeChannelMessageRemove);
  initializer.registerRtBefore("MatchCreate", beforeMatchCreate);
  initializer.registerRtBefore("MatchDataSend", beforeMatchDataSend);
  initializer.registerRtBefore("MatchJoin", beforeMatchJoin);
  initializer.registerRtBefore("MatchLeave", beforeMatchLeave);
  initializer.registerRtBefore("MatchmakerAdd", beforeMatchmakerAdd);
  initializer.registerRtBefore("MatchmakerRemove", beforeMatchmakerRemove);
  initializer.registerRtBefore("StatusFollow", beforeStatusFollow);
  initializer.registerRtBefore("StatusUnfollow", beforeStatusUnfollow);
  initializer.registerRtBefore("StatusUpdate", beforeStatusUpdate);
  initializer.registerRtBefore("Ping", beforePing);
  initializer.registerRtBefore("Pong", beforePong);


  initializer.registerLeaderboardReset(registerLeaderboardReset);
  //registerMatch < State = MatchState > (name: string, functions: MatchHandler<State>): void;
  initializer.registerMatch("", match);
  initializer.registerMatchmakerMatched(matchmakerMatched);
  initializer.registerTournamentEnd(tournamentEnd);
  initializer.registerTournamentReset(tournamentReset);

  initializer.registerRpc("healthcheck", rpcHealthCheck);

  nk.httpRequest(
    "http://redpanda:8082/topics/nakama",
    "post",
    {
      "Content-Type": "application/vnd.kafka.json.v2+json",
    },
    JSON.stringify({
      records: [{
        key: "nakama", value: {
          message: "hi", user: "nakama"
        }
      }]
    })
  );
};
