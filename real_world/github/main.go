package main

// go get golang.org/x/oauth2
// go get github.com/skratchdot/open-golang/open
import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

// Load from .env file using this library: "github.com/joho/godotenv"
func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	clientID = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
}

var (
	clientID     string
	clientSecret string
	redirectURL  = "http://localhost:18888"
	state        = "your state"
)

func main() {
	executeAPI()
}

func executeAPI() {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes: []string{
			"user:email",
			"public_repo",
			"gist",
			"read:user",
		},
		Endpoint: github.Endpoint,
	}
	// 初期化する
	var token *oauth2.Token

	file, err := os.Open("access_token.json")
	if os.IsNotExist(err) {
		// 初回アクセス
		url := conf.AuthCodeURL(state, oauth2.AccessTypeOnline)

		// コールバックを受け取るサーバー
		code := make(chan string)
		var server *http.Server
		server = &http.Server{
			Addr: ":18888",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// クエリパラメータから code を取得
				w.Header().Set("Content-Type", "text/html")
				io.WriteString(w, "<html><script>window.open('about:blank','_self').close()</script></html>")
				w.(http.Flusher).Flush()
				code <- r.URL.Query().Get("code")
				server.Shutdown(context.Background())
			}),
		}
		go server.ListenAndServe()

		open.Start(url)

		token, err := conf.Exchange(oauth2.NoContext, <-code)
		if err != nil {
			panic(err)
		}

		file, err := os.Create("access_token.json")
		if err != nil {
			panic(err)
		}
		json.NewEncoder(file).Encode(token)
	} else if err == nil {
		// 一度許可してローカルに保存済み
		token = &oauth2.Token{}
		fmt.Printf("token: %v", token)
		json.NewDecoder(file).Decode(token)
	} else {
		panic(err)
	}

	// header := http.Header{}
	// header.Set("Accept", "application/vnd.github.v3+json")
	// token.SetAuthHeader(&http.Request{
	// 	Header: header,
	// })
	client := oauth2.NewClient(oauth2.NoContext, conf.TokenSource(oauth2.NoContext, token))
	// Do Something
	// client.Get()
	fmt.Println("client is created!")

	getEmail(client)
}

func getEmail(client *http.Client) {
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	emails, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(emails))
}

func postGist(client *http.Client) {
	gist := `{
		"description": "API example",
		"public": true,
		"files": {
			"hello_from_rest_api.txt": {
				"content": "Hello world"
			}
		}
	}`

	// 投稿
	resp2, err := client.Post("https://api.github.com/gists", "application/json", strings.NewReader(gist))
	if err != nil {
		panic(err)
	}
	fmt.Println(resp2.Status)
	defer resp2.Body.Close()

	type GistResult struct {
		Url string `json:"html_url"`
	}

	gistResult := &GistResult{}
	// body, err := ioutil.ReadAll(resp2.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(body))
	err = json.NewDecoder(resp2.Body).Decode(&gistResult)
	if err != nil {
		panic(err)
	}
	if gistResult.Url != "" {
		open.Start(gistResult.Url)
	}
}
