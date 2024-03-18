package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// 何の後処理をするために closer を使うのか？
func main() {
	// // Close しなければならないものに、共通して当てはまることってなんだろう？
	// // no-op (= no operation) Closer
	// // io.NopCloser()

	// res, _ := http.Get("http://localhost:8080")
	// // res.Body は io.ReadCloser なので, Close が必要なのはわかる。
	// defer res.Body.Close()

	// var db sql.DB
	// rows, _ := db.Query("SELECT * FROM users")
	// // rows は *sql.Rows , これは？
	// defer rows.Close()

	// // io.Closer を満たしているようだ。
	// var _ io.Closer = rows

	// f, _ := os.Open("file")
	// defer f.Close()

	// ================= net/http =================
	// daareesa := "https://drive.usercontent.google.com/download?id=10SuuxuR_d47xZxDNOZmABrjDoKfFa4cM&export=download&authuser=0"
	// driveURL := "https://drive.google.com/file/d/10SuuxuR_d47xZxDNOZmABrjDoKfFa4cM/view?usp=sharing"
	// res, err := http.Get(driveURL)
	// if err != nil {
	// 	panic(err)
	// }

	// defer res.Body.Close()

	// ================= file =================
	f, _ := os.Create("file")
	defer func() {
		fmt.Printf("f.Sync(): %v\n", f.Sync())
		if ce := f.Close(); ce != nil {
			fmt.Printf("ce: %v\n", ce)
		}
		fmt.Printf("f.Sync(): %v\n", f.Sync())
	}()

	// この間に rm file するが、エラーにならない。。。
	fmt.Printf("f.Sync(): %v\n", f.Sync())
	time.Sleep(5 * time.Second)
	n, err := io.Copy(f, bytes.NewReader([]byte("hello, world")))
	fmt.Printf("n: %v\n", n)
	fmt.Printf("err: %v\n", err)

	mux := http.NewServeMux()
	// GETを指定したハンドラの登録
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}
