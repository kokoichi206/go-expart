const es = new EventSource("http://localhost:8192/stream");
es.addEventListener("message", function (event) {
  console.log(event.data);
  console.log(event);
});
es.addEventListener("error", function (event) {
  console.log(event);
  console.log(event.type);
  // 2: CLOSED, 1:  OPEN, 0: CONNECTING
  console.log(event.target.readyState);

  console.log("error: error: error: error: error: error piyopiyo tasuketeeee");
  es.close();
});
es.addEventListener("close", function (event) {
  console.log(event);
  console.log(event.type);
  // 2: CLOSED, 1:  OPEN, 0: CONNECTING
  console.log(event.target.readyState);

  console.log("closed.");
  es.close();
});
