package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Cat struct {
	Name string `cat:"name,required"`
	Age  int    `cat:"age"`
}

func main() {
	c := Cat{
		Name: "Tom",
		Age:  3,
	}
	v := reflect.ValueOf(c)

	fmt.Printf("v.Type(): %v\n", v.Type())

	t := v.Type()

	switch t.Kind() {
	case reflect.Struct:
		fmt.Println("v is a struct")
		parseCat(t)
	}
}

func parseCat(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field %d: %v\n", i, t.Field(i))

		sf := t.Field(i)
		// 埋め込みフィールドかどうか。
		if sf.Anonymous {
			t := sf.Type
			if t.Kind() == reflect.Pointer {
				t = t.Elem()
			}
			if !sf.IsExported() && t.Kind() != reflect.Struct {
				continue
			}
		} else if !sf.IsExported() {
			// 非公開フィールドはスキップ。
			continue
		}
		tag := sf.Tag.Get("cat")
		if tag == "-" {
			continue
		}
		tagName, opts := parseTag(tag)
		fmt.Printf("tagName: %v\n", tagName)
		fmt.Printf("opts: %v\n", opts)
	}
}

type tagOptions string

// 1つ目のタグを必須、それ以降をオプションとして扱う。
// 例) a,b,c => 'a' + 'b,c'
func parseTag(tag string) (string, tagOptions) {
	tag, opt, _ := strings.Cut(tag, ",")
	return tag, tagOptions(opt)
}

// options に関しては、個別で扱うことはせずに『含む』or『含まない』のみを扱う。
func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}
	s := string(o)
	for s != "" {
		var name string
		name, s, _ = strings.Cut(s, ",")
		if name == optionName {
			return true
		}
	}
	return false
}
