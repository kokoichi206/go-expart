## Part 1

```go
func main() {

    // /p <- Greedy matching
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbye World")
	})

	http.ListenAndServe(":9090", nil)
}
```

- [ResponseWriter](https://pkg.go.dev/net/http#ResponseWriter)
- [Http#Error()](https://pkg.go.dev/net/http#Error)

```sh
curl localhost:9090/goodbye

curl -v -d 'HEEEI' localhost:9090
```

http パッケージをよく見る！！！

```go
http.Error(w, "Ooops", http.StatusBadRequest)
w.WriteHeader(http.StatusBadRequest)
w.Write([]byte("Oooops"))
```

## Part 2

reference との使い分けがいまいち分かってないからかもしれんけど、なんで w は実態で r は参照なんだっけ

```
w http.ResponseWriter, r *http.Request
```

- タイムアウトを設けるのは、良いプラクティス！
- https://pkg.go.dev/net/http#Server
- DNS, handshake などのコストを下げるため、コネクションを繋いだままにしたりする
  - IdleTimeout
- os.signal#Notify: https://pkg.go.dev/os/signal#Notify

## Part 3

REST

- **struct tag** を使って Json 用のアノテーションをつけたりする！
  - `"-"` で無視することになる！
  - https://pkg.go.dev/encoding/json
- [REST by microsoft](https://learn.microsoft.com/en-us/azure/architecture/best-practices/api-design)

```sh
curl -v localhost:9090 -XDELETE
```

## Part 4

Standard Library offers a lot.

- encoding, decoding のロジックを data に持ってくるのが好き
- `%#v` で綺麗な出力を行う！
  - `p.l.Printf("Prod: %#v", prod)`

```
curl -v localhost:9090 -d '{"id": 1, "name": "tea"}'

curl -v localhost:9090/1 -XPUT
```

## Part 5

- gorilla
  - https://github.com/gorilla/mux
  - archived 2022/12/10 !!

## Part 6

- SKU is a object id, which is unique

## Part 7

- goswagger?
  - https://goswagger.io/
- meta data
  - https://goswagger.io/use/spec/meta.html

## English

- dig into
- Never mind
- Where do we go from here
