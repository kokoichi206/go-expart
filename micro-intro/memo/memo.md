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
