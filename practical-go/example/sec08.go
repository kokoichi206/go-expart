package example

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type account struct {
	BaseURL string `json:"baseURL"`
	APIKey  string `json:"apiKey"`
}

type Bottle struct {
	First      string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	Last       string `json:"last_name"`
	X          int    `json:"x,omitempty"`
	Y          *int   `json:"y,omitempty"`
	z          int    `json:"z,omitempty"` // export が必要
}

func ReadJson() {
	home, _ := os.UserHomeDir()
	fmt.Println("home", home)
	joined := filepath.Join(home, "/ghq/github.com/kokoichi206/go-expart/practical-go/example/test.json")
	fmt.Println(joined)
	f, err := os.Open(joined)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var resp account
	if err := json.NewDecoder(f).Decode(&resp); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", resp)

	b := Bottle{
		First: "First",
		Last:  "Last",
		X:     0, // 0 は出力されない！omitempty の対象
		Y:     Int(0),
	}
	out, _ := json.Marshal(b)
	fmt.Println(string(out))
}

func Int(v int) *int {
	return &v
}
