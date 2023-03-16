package main

import (
	"fmt"
	"time"
)

func channelTimer() {
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	fmt.Println("start timer")
	<-after
	fmt.Println("done")
}

func tickerTimer() {
	fmt.Println("5 seconds loop!!!")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now: ", now)
	}
}

func formatTime() {
	now := time.Now()
	fmt.Println(now.Format(time.RFC822))
	// MST time zone
	// PM 午前/午後
	// Z07 UTC との時差
	fmt.Println(now.Format("2006/01/02 03:04:05 MST"))
}