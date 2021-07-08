const MessageType = {
  Send: "send",
  Recv: "recv",
  Disconnect: "disc",
  Connected: "connect",
};

class MQTT {
  constructor(host, port, clientID, topic, msgHandler) {
    this.msgHandler = msgHandler;
    this.prepare(host, port, clientID, topic);
  }

  reconnect(host, port, clientID, topic) {
    if (this.client !== null && this.client !== undefined) {
      this.disconnect();
      this.client = null;
    }

    this.prepare(host, port, clientID, topic);
  }

  prepare(host, port, clientID, topic) {
    this.topic = topic;
    this.client = new Paho.MQTT.Client(host, Number(port), clientID);
    this.client.onConnectionLost = (res) => this.onConnectionLost(res);
    this.client.onMessageArrived = (message) => this.receive(message);
    this.client.connect({ onSuccess: () => this.onConnect() });
  }

  onConnect() {
    this.client.subscribe(this.topic);
    this.msgHandler.show(MessageType.Connected, "");
  }

  onConnectionLost(response) {
    if (response.errorCode !== 0) {
      console.error(`onConnectionLost: ${response.errorMessage}`);
    }
    this.msgHandler.show(MessageType.Disconnect, "");
  }

  disconnect() {
    this.client.disconnect();
    this.msgHandler.show(MessageType.Disconnect, "");
  }

  send(message) {
    const messageObj = new Paho.MQTT.Message(message);
    messageObj.destinationName = this.topic;
    this.client.send(messageObj);
    this.msgHandler.show(MessageType.Send, message);
  }

  receive(message) {
    console.info(`got message: ${message}`);
    this.msgHandler.show(MessageType.Recv, message.payloadString);
  }
}
