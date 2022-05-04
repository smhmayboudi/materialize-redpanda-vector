import { Kafka } from "kafkajs";

const kafka = new Kafka({
  clientId: "chat-app",
  brokers: ["localhost:9092"],
});

const producer = kafka.producer();

export function getConnection(user) {
  return producer.connect().then(() => {
    return (message) => {
      return producer.send({
        topic: "chat-room",
        messages: [{ value: JSON.stringify({ message, user }) }],
      });
    };
  });
}

export function disconnect() {
  return producer.disconnect();
}
