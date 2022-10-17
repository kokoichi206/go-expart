package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/kokoichi206/go-expert/web/todo/entity"
	"github.com/kokoichi206/go-expert/web/todo/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)

	// sut って何の略だったっけ
	sut := &KVS{Cli: cli}
	key := "Test_KVS_save"
	uid := entity.UserID(1234)
	ctx := context.Background()
	t.Cleanup(func() {
		// へー、Redis のメソッドにコメントないのか
		cli.Del(ctx, key)
	})
	if err := sut.Save(ctx, key, uid); err != nil {
		t.Errorf("want no error, but got %v", err)
	}
}

func TestKVS_Load(t *testing.T) {
	// どういう時に parallel にしていいかは整理したい
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)
	sut := &KVS{Cli: cli}

	// メソッド内で t.Run を呼ぶことで、サブテスト的に実行している
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		key := "Test_KVS_load_ok"
		uid := entity.UserID(1234)
		ctx := context.Background()
		// Duration = 30*time.Minute
		cli.Set(ctx, key, int64(uid), 30*time.Minute)
		t.Cleanup(func() {
			cli.Del(ctx, key)
		})

		got, err := sut.Load(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}
		if got != uid {
			// want, got っていいな。
			t.Errorf("want %d, but got %d", uid, got)
		}
	})

	t.Run("not Found", func(t *testing.T) {
		t.Parallel()

		key := "Test KVS_Save_notFound"
		ctx := context.Background()
		got, err := sut.Load(ctx, key)
		// ErrNotFound が返ってくる事が期待値！
		if err == nil || !errors.Is(err, ErrNotFound) {
			t.Errorf("want %v, but got %v (value = %d)", ErrNotFound, err, got)
		}
	})
}
