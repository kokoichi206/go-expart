package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	OccurredAt time.Time
}

func timeTimeV1() {
	// ウォール時間とモノトニック時間が含まれている
	t := time.Now()
	e1 := Event{OccurredAt: t}

	b, _ := json.Marshal(e1)

	var e2 Event
	// Unmarshal された time.Time にはウォール時間のみが含まれている！！
	_ = json.Unmarshal(b, &e2)

	// false
	fmt.Println(e1 == e2)
}

func timeTimeV2() {
	// ウォール時間とモノトニック時間が含まれている
	t := time.Now()
	e1 := Event{
		// モノトニック時間を切り捨てる！！
		OccurredAt: t.Truncate(0),
	}

	b, _ := json.Marshal(e1)

	var e2 Event
	// Unmarshal された time.Time にはウォール時間のみが含まれている！！
	_ = json.Unmarshal(b, &e2)

	// false
	fmt.Println(e1 == e2)
}
