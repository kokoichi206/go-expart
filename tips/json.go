package main

import (
	"encoding/json"
	"reflect"
	"unsafe"
)

type maskedValue struct {
	Value string
	mask  bool
}

func (v maskedValue) String() string {
	if v.mask {
		return "*"
	}
	return v.Value
}

type maskedConfig struct {
	Key   string
	Value maskedValue
}

type config struct {
	Key   string
	Value string
	mask  bool
}

func (c config) MarshalJSON() ([]byte, error) {
	value := c.Value
	if c.mask {
		for i := 0; i < len(c.Value); i++ {
			value = "*"
		}
	}
	// Q. なぜかここで panic が起きる。
	// A. 再帰的に MarshalJSON() が呼ばれるため。
	return json.Marshal(map[string]interface{}{
		"key":   c.Key,
		"value": value,
	})
}

func maskTest() {
	c := &config{
		Key:   "foo",
		Value: "bar",
		mask:  true,
	}
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	println(string(b))

	c2 := maskedConfig{
		Key: "foo",
		Value: maskedValue{
			Value: "bar",
			mask:  true,
		},
	}
	b2, err := json.Marshal(c2)
	if err != nil {
		panic(err)
	}
	println(string(b2))
}

type Me struct {
	name string
	age  int
}

var _ json.Marshaler = (*Me)(nil)
var _ json.Unmarshaler = (*Me)(nil)

func (m *Me) MarshalJSON() ([]byte, error) {
	// return json.Marshal(map[string]interface{}{
	// 	"name": m.name,
	// 	"age":  m.age,
	// })
	return s2b(`{"piyo": "pao"}`), nil
}

func (m *Me) UnmarshalJSON(b []byte) error {
	var tmp struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}

	m.name = tmp.Name
	m.age = tmp.Age
	return nil
}

func jsonTest() {
	m := &Me{
		name: "foo",
		age:  20,
	}

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	println(string(b))

	var m2 Me
	if err := json.Unmarshal(b, &m2); err != nil {
		panic(err)
	}

	println(m2.name, m2.age)
}

func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := *(*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Len = sh.Len
	bh.Cap = sh.Len
	return b
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
