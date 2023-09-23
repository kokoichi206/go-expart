package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func checkDBStatus(db *sql.DB) {
	stats := db.Stats()
	fmt.Printf("stats.Idle: %v\t", stats.Idle)
	fmt.Printf("stats.InUse: %v\n", stats.InUse)
}

func main() {
	source := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		"localhost", "4646", "ubuntu", "ubuntu", "testdb", "disable",
	)

	db, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal(err)
	}

	// sql.Open は実際にはコネクションをオープンしない！！
	checkDBStatus(db)

	// オープンするコネクションの最大数を設定する。
	// デフォルトでは 0 で無制限！！
	// db.SetMaxOpenConns(1)

	// 1 秒間アイドル状態が続いたコネクションは閉じる！
	// デフォルトでは 0 で無制限！！
	// db.SetConnMaxIdleTime(1 * time.Hour)

	defer func() {
		if closeErr := db.Close(); err != nil {
			log.Print(closeErr)
		}
	}()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Ping 等で、必要になった場合にコネクションをオープンする。
	checkDBStatus(db)

	ctx := context.Background()

	started := make(chan struct{})
	go func() {
		started <- struct{}{}
		db.ExecContext(ctx, "SELECT pg_sleep(10);")
	}()
	<-started
	checkDBStatus(db)

	go func() {
		started <- struct{}{}
		db.ExecContext(ctx, "SELECT pg_sleep(10);")
	}()
	<-started
	// Idle なコネクションがなく Max に達してない場合は、新たなコネクションがオープンされる。
	checkDBStatus(db)

	// デフォルトではコネクションは自発的に Close しない。
	// （ConnMaxIdleTime が設定されてないため。）
	time.Sleep(40 * time.Second)
}
