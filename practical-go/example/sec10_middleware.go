package example

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("start %s\n", r.URL)
		next.ServeHTTP(w, r)
		log.Printf("finish %s\n", r.URL)
	})
}

func middleware() {
	http.Handle("/healthz", MiddlewareLogging(http.HandlerFunc(Healthz)))
	http.ListenAndServe(":1344", nil)
}

// ミドルウェアでステータスコード等の情報を出力するため
//
// キャプチャするための構造体
type loggingResponseWriter struct {
	// 埋め込み
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	// statusCode を保存する！
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func wrapHandlerWithLogging(wrappedHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := NewLoggingResponseWriter(w)
		wrappedHandler.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		log.Printf("statusCode: %d, %s", statusCode, http.StatusText(statusCode))
	})
}

// ResponseWriter の Write メソッドを拡張する
func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	if lrw.statusCode >= 400 {
		//
		log.Printf("ERROR %d: Response Body: %s", lrw.statusCode, b)
	}
	return lrw.ResponseWriter.Write(b)
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover により、ハンドラーで発生した panic から復帰！
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// TX かいし、ロールバック等はミドルウェアで扱いたいが、
// 実際の処理は HandlerFunc で扱うので WithContext に乗せて運ぶ
func NewMiddlewareTx(db *sql.DB) func(http.Handler) http.Handler {
	return func(wrappedHandler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tx, _ := db.Begin()
			lrw := NewLoggingResponseWriter(w)
			r = r.WithContext(context.WithValue(r.Context(), "tx", tx))

			wrappedHandler.ServeHTTP(lrw, r)

			statusCode := lrw.statusCode
			if 200 <= statusCode && statusCode < 400 {
				log.Println("transaction committed")
				tx.Commit()
			} else {
				log.Print("rollback")
				tx.Rollback()
			}
		})
	}
}

func extractTx(r *http.Request) *sql.Tx {
	tx, ok := r.Context().Value("tx").(*sql.Tx)
	if !ok {
		panic("transaction middleware is not supported")
	}
	return tx
}

func txTest() {
	db := &sql.DB{}
	tx := NewMiddlewareTx(db)

	http.Handle("/comments", tx(Recovery(http.HandlerFunc(Comments))))
}

func Comments(w http.ResponseWriter, r *http.Request) {
	tx := extractTx(r)
	// DB アクセス処理！
	fmt.Println(tx)
}
