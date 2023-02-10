package main

import (
	"encoding/json"
	"fmt"
)

func ListData() ([]BigStruct, error) {
	res := []BigStruct{}

	err := listData(func(body []byte) error {
		r := BigStruct{}
		if err := json.Unmarshal(body, &r); err != nil {
			return fmt.Errorf("failed to unmarshal json: %w", err)
		}

		res = append(res, r)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to listData: %w", err)
	}

	return res, nil
}

func listData(logic func(body []byte) error) error {
	body := []byte(`{"member":"test-member", "info": [{"name":"test-name"}]}`)

	for j := 0; j < 10; j++ {
		for i := 0; i < 1000; i++ {
			if err := logic(body); err != nil {
				return fmt.Errorf("failed to execute logic: %w", err)
			}
		}
	}

	return nil
}

type BigStruct struct {
	Member *string `json:"member"`
	Info   []Info  `json:"info"`
}

type Info struct {
	Name string `json:"name"`
}
