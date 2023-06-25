package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	query, err := url.QueryUnescape("2023-06-21T12%3A23%3A11.938%2B09%3A00")
	fmt.Printf("err: %v\n", err)
	fmt.Printf("query: %v\n", query)

	rawQuery := "last_read_at=2023-06-21T12:23:11.938+09:00"
	values, err := url.ParseQuery(rawQuery)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("values: %v\n", values)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("r.URL.Query(): %v\n", r.URL.Query())
		fmt.Printf("r.URL.Query()[\"calc\"]: %v\n", r.URL.Query()["calc"])
		fmt.Printf("r.URL.RawQuery: %v\n", r.URL.RawQuery)
	})

	http.ListenAndServe(":8980", nil)
}
