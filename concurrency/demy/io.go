package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

func dbSample() {
	db, _ := sql.Open("mysql", "test_user:test_password@tcp(127.0.0.1:13306)/test_db")
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Dolly")
}

func httpCall() {
	req, _ := http.NewRequest("GET", "http://example.com", nil)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	req = req.WithContext(ctx)

	client := http.Client{
		Timeout: time.Second,
	}

	res, _ := client.Do(req)
	fmt.Printf("res: %v\n", res)
}
