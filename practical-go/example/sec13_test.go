package example

import (
	"bytes"
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSec13(t *testing.T) {
	type args struct {
		a        int
		b        int
		operator string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Add",
			args: args{
				a:        10,
				b:        2,
				operator: "+",
			},
			want:    12,
			wantErr: false,
		},
		{
			name: "Invalid Operator",
			args: args{
				a:        10,
				b:        2,
				operator: "?",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		// t.Parallel に必要
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := Calc(tt.args.a, tt.args.b, tt.args.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("Calc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Calc(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	default:
		return -99999999, errors.New("Unknown operator")
	}
}

// TestMain に事前・事後処理を書く！
func TestMain(m *testing.M) {
	// 事前処理
	setup()

	// 事後処理
	defer teardown()

	m.Run()
}

func setup()    {}
func teardown() {}

func TestCmp(t *testing.T) {
	type user struct {
		first  string
		last   string
		age    int
		skills []string
	}
	u1 := user{
		first:  "John",
		last:   "Doe",
		age:    320,
		skills: []string{"go", "Kotlin"},
	}
	u2 := user{
		first:  "John",
		last:   "Doe",
		age:    320,
		skills: []string{"go", "Kotlin"},
	}

	// export されてないフィールドを比較対象にするオプション
	opt := cmp.AllowUnexported(user{})
	if diff := cmp.Diff(u1, u2, opt); diff != "" {
		t.Errorf("User value is mismatch (-u1 + u2):\n%s", diff)
	}
}

func TestConsoleOut(t *testing.T) {
	var b bytes.Buffer
	DumpUserTo(&b, &User{Name: "John Doe"})
	if b.String() != "John Doe(住所不定)" {
		t.Errorf("error (expected: '...~~', actual='%s'", b.String())
	}
}
