package example

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type txAdmin struct {
	*sql.DB
}

type Service struct {
	tx txAdmin
}

// トランザクションを制御するメソッド
func (t *txAdmin) Transaction(ctx context.Context, f func(ctx context.Context) (err error)) error {
	tx, err := t.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// 確実に rollback する。
	defer tx.Rollback()

	if err := f(ctx); err != nil {
		return fmt.Errorf("trancation query failed: %w", err)
	}
	return tx.Commit()
}

func (s *Service) UpdateProduct(ctx context.Context, productId string) error {
	updateFunc := func(ctx context.Context) error {
		if _, err := s.tx.ExecContext(ctx, "SELECT * FROm ..."); err != nil {
			return err
		}
		if _, err := s.tx.ExecContext(ctx, "UPDATE ..."); err != nil {
			return err
		}
		return nil
	}
	return s.tx.Transaction(ctx, updateFunc)
}

func (t *txAdmin) Tuning() {
	t.DB.SetConnMaxIdleTime(3 * time.Hour)
	t.DB.SetMaxIdleConns(10)
}

func (t *txAdmin) CancellableQuery() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	if _, err := t.ExecContext(ctx, "SELECT pg_sleep(100)"); err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("canceling query")
		} else {
			// Unexpected error...
		}
	}
}
