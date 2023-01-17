package example

import (
	"fmt"
	"reflect"
	"strconv"
)

type Author struct {
	FirstName string
	LastName  string
}

type Book struct {
	Title  string
	Author Author
	ISBN   string
}

func NewAuthor(first, last string) *Author {
	return &Author{
		FirstName: first,
		LastName:  last,
	}
}

func test() {
	_ = NewAuthor("a", "b")
	a := Author{
		FirstName: "a",
		LastName:  "b",
	}
	book := Book{
		Title:  "a",
		Author: a,
		ISBN:   "c",
	}
	fmt.Println("Title: ", book.Title)
}

// p66 ~
// Reflect パッケージを使ったタグの取り扱い
type MapStruct struct {
	Str    string  `map:"str"`
	StrPtr *string `map:"str"`
	Bool   bool    `map:"bool"`
}

func TagTest() {
	src := map[string]string{
		"str":  "string data",
		"bool": "true",
		"int":  "123459",
	}
	var ms MapStruct
	Decode(&ms, src)
	fmt.Println(ms)
}

func Decode(target interface{}, src map[string]string) error {
	v := reflect.ValueOf(target)
	e := v.Elem()
	return decode(e, src)
}

func decode(e reflect.Value, src map[string]string) error {
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		// 再帰処理
		if f.Anonymous {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// 子が構造体だったら再帰処理
		if f.Type.Kind() == reflect.Struct {
			if err := decode(e.Field(i), src); err != nil {
				return err
			}
			continue
		}

		// タグがない場合、フィールド名をそのまま使う
		key := f.Tag.Get("map")
		if key == "" {
			key = f.Name
		}

		// 元データにない場合
		sv, ok := src[key]
		if !ok {
			continue
		}

		// フィールドの型を取得
		var k reflect.Kind
		var isP bool
		if f.Type.Kind() != reflect.Ptr {
			k = f.Type.Kind()
		} else {
			k = f.Type.Elem().Kind()

			// ポインターのポインター
			if k == reflect.Ptr {
				continue
			}
			isP = true
		}

		switch k {
		case reflect.String:
			if isP {
				e.Field(i).Set(reflect.ValueOf(&sv))
			} else {
				e.Field(i).SetString(sv)
			}
		case reflect.Bool:
			b, err := strconv.ParseBool(sv)
			if err == nil {
				if isP {
					e.Field(i).Set(reflect.ValueOf(&b))
				} else {
					e.Field(i).SetBool(b)
				}
			}
		case reflect.Int:
			n64, err := strconv.ParseInt(sv, 10, 64)
			if err == nil {
				if isP {
					n := int(n64)
					e.Field(i).Set(reflect.ValueOf(&n))
				} else {
					e.Field(i).SetInt(n64)
				}
			}
		}
	}
	return nil
}
