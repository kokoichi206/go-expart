let socket = new WebSocket("ws://localhost:3333/ws")
socket.onmessage = (event) => { console.log("received from the server: ", event.data) }
(event) => { console.log("received from the server: ", event.data) }
socket.send("hello from client")

let socket = new WebSocket("ws://localhost:3333/orderbook-feed")
socket.onmessage = (event) => { console.log("received from the server: ", event.data) }
