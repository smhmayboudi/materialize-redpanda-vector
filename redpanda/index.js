import readline from "readline";

import * as Producer from "./producer.js";
import * as Consumer from "./consumer.js";

const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
});

function start() {
  console.log("connecting...");
  Consumer.connect().then(() => {
    rl.question("enter user name \n", function (username) {
      Producer.getConnection(username).then((sendMessage) => {
        console.log("connected, press Ctrl+C to exit");
        rl.on("line", (input) => {
          readline.moveCursor(process.stdout, 0, -1);
          sendMessage(input);
        });
      });
    });
  });
}

start();

process.on("SIGINT", process.exit);

process.on("exit", () => {
  Producer.disconnect();
  Consumer.disconnect();
  rl.close();
});
