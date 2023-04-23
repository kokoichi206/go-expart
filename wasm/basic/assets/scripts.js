const goWasm = new Go();

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance);

        document.getElementById("get-html-button")
            .addEventListener("click", () => {
                document.body.innerHTML += getHtml();
            })
    });
 