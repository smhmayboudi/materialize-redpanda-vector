import { v4 as uuidv4 } from "uuid";
import { Kafka } from "kafkajs";

const kafka = new Kafka({
  clientId: "chat-app",
  brokers: ["localhost:9092"],
});

const consumer = kafka.consumer({ groupId: uuidv4() });

export function connect() {
  return consumer.connect().then(() =>
    consumer.subscribe({ topic: "chat-room" }).then(() =>
      consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
          const formattedValue = JSON.parse(message.value.toString());
          console.log(`${formattedValue.user}: ${formattedValue.message}`);
        },
      })
    )
  );
}

export function disconnect() {
  consumer.disconnect();
}
