package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"
)

func dbSample() {
	db, _ := sql.Open("mysql", "test_user:test_password@tcp(127.0.0.1:13306)/test_db")
	defer db.Close()

	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	db.ExecContext(ctx, "INSERT INTO users (name) VALUES (?)", "Dolly")
	db.QueryContext(ctx, "SELECT * FROM users")
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

func cmd() {
	fmt.Println("========== cmd ==========")
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sleep", "3")
	err := cmd.Run()

	fmt.Printf("err: %v\n", err)
	fmt.Printf("errors.Is(err, os.ErrDeadlineExceeded): %v\n", errors.Is(err, os.ErrDeadlineExceeded))
}

func cmdSample() {
	fmt.Println("========== cmdSample ==========")
	// ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// cmd := exec.CommandContext(ctx, "sleep", "3")
	cmd := exec.CommandContext(ctx, "echo", "hoge", ">", "/dev/null")
	cmd.Start()
	fmt.Printf("cmd.ExtraFiles: %v\n", cmd.ExtraFiles)

	err := cmd.Run()
	fmt.Printf("err: %v\n", err)
}
