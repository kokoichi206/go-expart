package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func currentGoroutine() {
	for range time.Tick(time.Second * 10) {
		fmt.Printf("runtime.NumGoroutine(): %v\n", runtime.NumGoroutine())
	}
}

func gingin() {
	go currentGoroutine()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	router.GET("/stream", addSseHeaders(), func(c *gin.Context) {
		for range time.Tick(time.Second * 3) {
			select {
			case <-c.Request.Context().Done():
				fmt.Println("client closed")
			default:
			}
			now := time.Now().Format("2006-01-02 15:04:05")
			currentTime := fmt.Sprintf("The Current Time Is %v", now)

			fmt.Println("sent")

			c.SSEvent("message", currentTime)

			c.Writer.Flush()
		}
	})

	router.Run(":8192")
}

func main() {
	// gingin()

	standard()
}

func addSseHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")
		c.Writer.Header().Set("Transfer-Encoding", "chunked")
		c.Writer.Header().Set("X-Accel-Buffering", "no")
		c.Next()
	}
}

func addSseHeadersMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("X-Accel-Buffering", "no")

		next.ServeHTTP(w, r)
	})
}

func standard() {
	mux := http.NewServeMux()

	// ミドルウェア。一番外側に当てたいものから追加する。
	baseMWs := []func(http.Handler) http.Handler{}
	baseMWs = append(baseMWs, recovery)
	baseMWs = append(baseMWs, logging)

	m := &maxMux{
		mux: mux,
		mws: baseMWs,
	}

	m.handle(http.MethodGet, "/recovery", failed)
	m.handle(http.MethodGet, "/hi", hi)
	// // 全角が入っていると 404 になる。
	// m.handle(http.MethodGet, "GET　/pien", hi)
	m.handle(http.MethodGet, "/ぴえん/{name}", pien)
	m.handle(http.MethodGet, "/json", jsonRes)

	// 以降のメソッドは、SSE 用のヘッダを mw で付与する。
	m.mws = append(m.mws, addSseHeadersMW)
	m.handle(http.MethodGet, "/stream", stream)

	dir, _ := os.Getwd()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/"))))

	m.run("8192")
}

type maxMux struct {
	mux      *http.ServeMux
	mws      []func(http.Handler) http.Handler
	patterns []string
}

func (m *maxMux) handle(method string, path string, fn func(w http.ResponseWriter, r *http.Request)) {
	handler := http.Handler(http.HandlerFunc(fn))

	// ミドルウェアを逆順に適用する。
	for idx := range m.mws {

		mw := m.mws[len(m.mws)-1-idx]
		handler = mw(handler)
	}

	pattern := fmt.Sprintf("%s %s", method, path)
	m.mux.Handle(pattern, handler)

	m.patterns = append(m.patterns, pattern)
}

func (m *maxMux) run(port string) {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: m.mux,
	}

	fmt.Printf("server listening on :%s\n", port)
	for _, p := range m.patterns {
		fmt.Printf("pattern: %v\n", p)
	}
	server.ListenAndServe()
}

func failed(w http.ResponseWriter, r *http.Request) {
	panic("failed")
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic recovered: (message = %s)\n", err)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("logging: {path: %v}\n", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

// http://localhost:8192/ぴえん/fuga?paon=テストなう
func pien(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pien\n"))
	w.Write([]byte(fmt.Sprintf("name: %v\n", r.PathValue("name"))))
	w.Write([]byte(fmt.Sprintf("paon: %v\n", r.URL.Query().Get("paon"))))
}

func stream(w http.ResponseWriter, r *http.Request) {
	for range time.Tick(time.Second * 3) {
		select {
		case <-r.Context().Done():
			fmt.Println("client closed")
		default:
		}
		now := time.Now().Format("2006-01-02 15:04:05")
		currentTime := fmt.Sprintf("The Current Time Is %v", now)

		fmt.Println("sent")

		eventType := "message"
		fmt.Fprintf(w, "event:%s\n", eventType)
		fmt.Fprintf(w, "data:%s\n\n", currentTime)
		w.(http.Flusher).Flush()
	}
}

type Me struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func jsonRes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	m := Me{
		Name: "John",
		Age:  30,
	}

	json.NewEncoder(w).Encode(m)
}
