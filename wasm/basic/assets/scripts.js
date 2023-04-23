const goWasm = new Go();

console.log("instantiateStreaming");

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject).then(
  (result) => {
    goWasm.run(result.instance);

    document.getElementById("get-html-button").addEventListener("click", () => {
      document.body.innerHTML += getHtml();
    });

    console.log("calculating");

    let x = add(1, 2);
    console.log("calculated");
    console.log(x);

    // TEST: error expected patterns
    // panic: syscall/js: call of Value.Int on string
    // let y = add(1, "pien");
    // panic: runtime error: index out of range [1] with length 1
    // y = add(1);
  }
);
