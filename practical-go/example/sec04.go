package example

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gen2brain/beeep"
)

type Notifier interface {
	Show(message string)
}

type ConsoleWarning struct{}

func (c ConsoleWarning) Show(message string) {
	fmt.Fprintf(os.Stderr, "[%s]: %s\n", os.Args[0], message)
}

type DesktopWarning struct{}

func (d DesktopWarning) Show(message string) {
	beeep.Alert(os.Args[0], message, "")
}

type SlackWarning struct {
	URL     string
	Channel string
}

type SlackMessage struct {
	Text      string `json:"text"`
	Username  string `json:"username"`
	IconEmoji string `json:"icon_emoji"`
	Channel   string `json:"channel"`
}

func (s SlackWarning) Show(mesasge string) {
	params, _ := json.Marshal(SlackMessage{
		Text:      mesasge,
		Username:  os.Args[0],
		IconEmoji: ":robot_face:",
		Channel:   s.Channel,
	})

	resp, err := http.PostForm(
		s.URL,
		url.Values{"payload": {string(params)}},
	)
	if err == nil {
		io.ReadAll(resp.Body)
		defer resp.Body.Close()
	}
}

func Parser(r io.Reader) {
	if c, ok := r.(io.Closer); ok {
		c.Close()
	}
}

func Interface() {
	var notifier Notifier

	notifier = &ConsoleWarning{}
	notifier.Show("Hello world to console")

	notifier = &DesktopWarning{}
	notifier.Show("Desktop notify!")

	// ダミーの変数宣言をすることで確かめることもできる。
	var _ Notifier = &SlackWarning{}

	// 型アサーション
	ctx := context.WithValue(context.Background(), "favorite", "ゲーム")
	if s, ok := ctx.Value("favorite").(string); ok {
		log.Printf("私の好きなものは%sです\n", s)
	}
}
