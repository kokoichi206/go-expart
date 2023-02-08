package example

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

func LogPackage() {
	file, _ := os.Create("log.txt")
	log.SetOutput(io.MultiWriter(file, os.Stderr))
	log.Println("Print to file and Standard Error")

	// Log.Llongfile
	log.SetFlags(log.LstdFlags)
	log.SetPrefix("🐸 ")
	log.Println("ueeei")
}

const DBConnErrCD = "E10001"

func businessLogic() {
	// log のなかに, "code": DBConnErrCD とかで埋め込む
}

func customHttpErrorLog() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world!")
	}
	http.HandleFunc("/hello", handler)

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Server のエラーログは標準出力に勝手に出てしまう
	// 今回はそれを書き換える。
	// （io.Writer を満たすロガー！）
	// io.Writer を満たさない場合は, Write(p []byte)(int, error) を実装することで
	// io.Writer を満たすようにする
	server := &http.Server{
		Addr: ":23333",
		// io.Writer, prefix, flag
		ErrorLog: log.New(logger, "", 0),
	}

	logger.Fatal().Msgf("server: %v", server.ListenAndServe())
}
