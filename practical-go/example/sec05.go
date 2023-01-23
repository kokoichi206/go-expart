package example

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var ErrNotFound = errors.New("NOT FOUND!")

func findBook(isbn string) (*Book, error) {
	return nil, ErrNotFound
}

func validate(length int) error {
	if length <= 0 {
		return fmt.Errorf("length must be greater than 0, length = %d", length)
	}

	return nil
}

type HTTPError struct {
	StatusCode int
	URL        string
}

// HTTPError のポインターが Error インタフェースを満たす！
func (h *HTTPError) Error() string {
	return fmt.Sprintf("http status code = %d, url =%s", h.StatusCode, h.URL)
}

func readContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Error インタフェースを満たす、HTTPError のポインターを error として返してあげる
		return nil, &HTTPError{
			StatusCode: resp.StatusCode,
			URL:        url,
		}
	}
	return io.ReadAll(resp.Body)
}

// エラーの wrap, unwrap
type loadConfigError struct {
	msg string
	err error
}

func (l *loadConfigError) Error() string {
	return fmt.Sprintf("cannot load config: %s (%s)", l.msg, l.err.Error())
}

func (l *loadConfigError) Unwrap() error {
	return l.err
}

type Config struct {
	Port string
	Host string
}

func LoadCOnfig(configFilePath string) (*Config, error) {
	var cfg *Config
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("read file `%s`", configFilePath),
			err: err,
		}
	}
	if err = json.Unmarshal(data, cfg); err != nil {
		return nil, &loadConfigError{
			msg: fmt.Sprintf("parse config file `%s`", configFilePath),
			err: err,
		}
	}
	return cfg, nil
}

func ErrorIs() {
	err := sql.ErrNoRows
	// true
	fmt.Println(errors.Is(err, sql.ErrNoRows))

	wrapped := fmt.Errorf("fail to get from db: %w", err)
	// true
	fmt.Println(errors.Is(wrapped, sql.ErrNoRows))
}

// `log.Fatal()` などを呼び出すとプログラムが即座に終了し
// defer でリソースがクローズされない！
func FatalCheck() {
	// Will not called !!
	defer func() {
		println("defer called 1")
	}()

	// Will not called !!
	defer func() {
		println("defer called 2")
	}()

	log.Fatal("something unexpected happened. Exit with 1.")

	// Will not called !!
	defer func() {
		println("defer called 3")
	}()
}
