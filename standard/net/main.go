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
	mux.HandleFunc("GET /hi", hi)
	// 全角が入っていると 404 になる。
	mux.HandleFunc("GET　/pien", hi)
	mux.HandleFunc("GET /ぴえん/{name}", pien)
	mux.Handle("/stream", addSseHeadersMW(http.HandlerFunc(stream)))
	mux.Handle("/json", addSseHeadersMW(http.HandlerFunc(jsonRes)))

	server := &http.Server{
		Addr:    ":8192",
		Handler: mux,
	}

	// go currentGoroutine()
	dir, _ := os.Getwd()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/static/")))) //追記

	server.ListenAndServe()
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
