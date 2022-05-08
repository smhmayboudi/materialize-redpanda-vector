CREATE SOURCE chat_room
FROM KAFKA BROKER 'redpanda:29092' TOPIC 'chat-room' FORMAT BYTES ENVELOPE NONE;
/*
 CREATE SOURCE "materialize"."public"."chat_room"
 FROM KAFKA BROKER 'redpanda:29092' TOPIC 'chat-room' FORMAT BYTES ENVELOPE NONE;
 */
CREATE MATERIALIZED VIEW chat_room_view AS
SELECT data->>'message' AS message,
  data->>'user' AS user
FROM (
    SELECT CONVERT_FROM(data, 'utf8')::jsonb AS data
    FROM chat_room
  );
/*
 CREATE VIEW "materialize"."public"."chat_room_view" AS
 SELECT "data"->>'message' AS "message",
 "data"->>'user' AS "user"
 FROM (
 SELECT "pg_catalog"."convert_from"("data", 'utf8')::"pg_catalog"."jsonb" AS "data"
 FROM "materialize"."public"."chat_room"
 );
 */
CREATE SINK chat_room_sink
FROM chat_room_view INTO KAFKA BROKER 'redpanda:29092' TOPIC 'chat-room-materialized' CONSISTENCY (
    TOPIC 'chat-room-materialized-consistency' FORMAT AVRO USING CONFLUENT SCHEMA REGISTRY 'http://redpanda:8081'
  ) WITH (reuse_topic = true) FORMAT JSON;
/*
 CREATE SINK "materialize"."public"."chat_room_sink" IN CLUSTER [1]
 FROM [u2 AS "materialize"."public"."chat_room_view"] INTO KAFKA BROKER 'redpanda:29092' TOPIC 'chat-room-materialized' CONSISTENCY (
 TOPIC 'chat-room-materialized-consistency' FORMAT AVRO USING CONFLUENT SCHEMA REGISTRY 'http://redpanda:8081'
 ) WITH ("reuse_topic" = true) FORMAT JSON WITH SNAPSHOT;
 */