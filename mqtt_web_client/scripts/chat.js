class ChatController {
  constructor() {
    this.chatbox = document.getElementById("chatbox");
  }

  show(msgType, message) {
    let payload;

    switch (msgType) {
      case MessageType.Send:
        payload = `<span class="sender-id">You: </span>${message}<br>`;
        break;
      case MessageType.Recv:
        payload = `<span class="receive-id">Stranger: </span>${message}<br>`;
        break;
      case MessageType.Connected:
        payload = `<span class="info-id">You're now chatting with a random stranger. Say hi!</span><br>`;
        break;
      case MessageType.Disconnect:
        payload = `<span class="info-id">You has been disconnected</span><br>`;
        break;
      default:
        console.error(`unrecognized message: ${msgType}`);
        return;
    }

    this.chatbox.innerHTML = this.chatbox.innerHTML + payload;
    this.chatbox.scrollTop =
      this.chatbox.scrollHeight - this.chatbox.clientHeight;
  }

  clear() {
    this.chatbox.innerHTML = "";
  }
}
