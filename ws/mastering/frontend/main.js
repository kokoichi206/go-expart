const selectedChat = "general";

class Event {
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

function routeEvent(event) {
  if (event.type == undefined) {
    alert("No type field in even");

    switch (event.type) {
      case "new_message":
        console.log("NEW Message!");
        break;
      default:
        alert("unsupported message type!");
        break;
    }
  }
}

function sendEvent(eventName, payload) {
  const event = new Event(eventName, payload);
  conn.send(JSON.stringify(event));
}

function changeChatRoom() {
  const newChat = document.getElementById("chatroom");
  if (newChat != null && newChat.value != selectedChat) {
    console.log(newChat);
  }
  return false;
}

function sendMessage() {
  const newMessage = document.getElementById("message");
  if (newMessage != null) {
    console.log(newMessage);
    // conn.send(newMessage.value);
    sendEvent("send_message", newMessage.value);
  }
  return false;
}
console.log("eeee...");

window.onload = function () {
  document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
  document.getElementById("chatroom-message").onsubmit = sendMessage;

  console.log("onload...");
  if (window["WebSocket"]) {
    console.log("supports ws!!");
    // connect to ws
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onmessage = function (evt) {
      // console.log(evt);
      const eventData = JSON.parse(eve.data);
      const event = Object.assign(new Event(), eventData);
      routeEvent(event);
    };
  } else {
    alert("browser doesn't support ws!");
  }
};
