// SQL 実行時に利用する時刻情報を制御するためのパッケージ
//
// インタフェースの抽象化レイヤーを用いることで、
// テスト用に固定値を使えるようにしている！
//
package clock

import (
	"time"
)

type Clocker interface {
	Now() time.Time
}

// アプリケーションで実際に使う
type RealClocker struct{}

func (r RealClocker) Now() time.Time {
	return time.Now()
}

// テスト用の固定時刻を返す
type FixedClocker struct{}

func (fc FixedClocker) Now() time.Time {
	return time.Date(2022, 10, 11, 14, 28, 10, 0, time.UTC)
}
