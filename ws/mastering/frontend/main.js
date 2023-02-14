const selectedChat = "general";

class Event {
  constructor(type, payload) {
    this.type = type;
    this.payload = payload;
  }
}

class SendMessageEvent {
  constructor(message, from) {
    this.message = message;
    this.from = from;
  }
}

class NewMessageEvent {
  constructor(message, from, sent) {
    this.message = message;
    this.from = from;
    this.sent = sent;
  }
}

function routeEvent(event) {
  if (event.type == undefined) {
    alert("No type field in even");
  }
  switch (event.type) {
    case "new_message":
      console.log("NEW Message!");
      const messageEvent = Object.assign(new NewMessageEvent(), event.payload);
      appendChatMessage(messageEvent);
      break;
    default:
      alert("unsupported message type!");
      break;
  }
}

function appendChatMessage(messageEvent) {
  const date = new Date(messageEvent.sent);
  console.log("date.toLocaleString: " + date.toLocaleString());
  const formattedMsg = `${date.toLocaleString()}: ${messageEvent.message}`;
  textarea = document.getElementById("chatmessages");
  textarea.innerHTML = textarea.innerHTML + "\n" + formattedMsg;
  textarea.scrollTop = textarea.scrollHeight;
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
    // message, username
    let outgoingEvent = new SendMessageEvent(newMessage.value, "kokoichi");
    console.log(newMessage);
    // conn.send(newMessage.value);
    sendEvent("send_message", outgoingEvent);
  }
  return false;
}
console.log("eeee...");

function login() {
  let formData = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
  };
  fetch("login", {
    method: "post",
    body: JSON.stringify(formData),
    mode: "cors",
  })
    .then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        throw "unauthorized!";
      }
    })
    .then((data) => {
      // now, we are authenticated!
      connectWebsocket(data.otp);
    })
    .catch((e) => {
      console.alert(e);
    });

  return false;
}

function connectWebsocket(otp) {
  console.log("connectWebsocket..." + otp);
  if (window["WebSocket"]) {
    console.log("supports ws!!");
    // connect to ws
    conn = new WebSocket("wss://" + document.location.host + "/ws?otp=" + otp);

    conn.onopen = function (evt) {
      document.getElementById("connection-header").innerHTML =
        "Connected to WebSocket: true";
    };
    conn.onclose = function (evt) {
      document.getElementById("connection-header").innerHTML =
        "Connected to WebSocket: false";
      // automatic reconnection is better
    };

    conn.onmessage = function (evt) {
      // console.log(evt);
      const eventData = JSON.parse(evt.data);
      const event = Object.assign(new Event(), eventData);
      routeEvent(event);
    };
  } else {
    alert("browser doesn't support ws!");
  }
}

window.onload = function () {
  document.getElementById("chatroom-selection").onsubmit = changeChatRoom;
  document.getElementById("chatroom-message").onsubmit = sendMessage;
  document.getElementById("login-form").onsubmit = login;
};
