package example

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

type Func func() string

func (f Func) String() string {
	return f()
}

type Comment struct {
	Message  string
	UserName string
}

func ServeTest() {
	var mutex = &sync.RWMutex{}
	comments := make([]Comment, 0, 100)

	http.HandleFunc("/comments", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			mutex.RLock() // 読み込み時に書き込みがあることを考慮しロックする

			if err := json.NewEncoder(w).Encode(comments); err != nil {
				// http.Error よい
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.RUnlock()

			// FormValue() を呼ぶ場合は省略可能
			_ = r.ParseForm()
			for key, values := range r.Form {
				// 同じキーが複数来た場合 values に複数格納される
				fmt.Printf(" %s: %v\n", key, values)
			}

		case http.MethodPost:
			var c Comment
			if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
				http.Error(w, fmt.Sprintf(`{"status":"%s"}`, err), http.StatusInternalServerError)
				return
			}
			mutex.Lock() // 同時に複数アクセスを防ぐためにロックする
			comments = append(comments, c)
			mutex.Unlock()

			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(`{"status":"created"}`))
		default:
			http.Error(w, `{"status":"only GET or POST methods are permitted"}`, http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8888", nil)
}

func chiMux() {
	yes := 0
	no := 0
	r := chi.NewRouter()
	r.Post("/poll/{answer}", func(w http.ResponseWriter, r *http.Request) {
		if chi.URLParam(r, "answer") == "y" {
			yes++
		} else {
			no++
		}
	})
	r.Get("/result", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "agree: %d, disagree: %d", yes, no)
	})
	r.Handle("/asset/*", http.StripPrefix("/asset/", http.FileServer(http.Dir("."))))

	log.Fatal(http.ListenAndServe(":8192", r))
}
