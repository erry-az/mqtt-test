let client = null;
let chatController = null;
let modal = null;

function toggleSend() {
  const { value } = document.getElementById("chat");
  const sendBtn = document.getElementById("send");
  sendBtn.disabled = value.length < 1;
}

function onSend() {
  const chatDOM = document.getElementById("chat");
  const { value: msg } = chatDOM;
  client.send(msg);
  chatDOM.value = "";
  toggleSend();
}

function newConnection() {
  const chatInput = document.getElementById("chat");
  chatInput.disabled = true;

  if (chatController == null) {
    chatController = new ChatController();
  }

  const host = document.getElementById("config-host").value;
  const port = document.getElementById("config-port").value;
  const clientID = document.getElementById("config-client-id").value;
  const topic = document.getElementById("config-topic").value;

  if (client == null) {
    client = new MQTT(host, port, clientID, topic, chatController);
  } else {
    client.reconnect(host, port, clientID, topic);
  }

  chatController.clear();
  toggleModal();
  showConfigClose(true);
  chatInput.disabled = false;
}

function showConfigClose(show) {
  const display = show ? "block" : "none";
  const closeButtons = document.getElementsByClassName("config-close");
  for (let idx = 0; idx < closeButtons.length; idx++) {
    closeButtons[idx].style.display = display;
  }
}

function toggleModal() {
  modal.toggle();
}

function main() {
  if (modal == null) {
    modal = new bootstrap.Modal(document.getElementById("config-modal"));
  }
  toggleModal();
}

window.onload = main;
document.addEventListener("keydown", (event) => {
  if (client === null) return;
  handleEsc(event);
});
