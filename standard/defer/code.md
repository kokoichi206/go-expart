``` sh
❯ go build main.go
❯ file main
main: Mach-O 64-bit executable arm64
```

環境

``` sh
$ go version
go version go1.20.7 linux/arm64
ubuntu@ubuntu:~/work/standard$ uname -a
Linux ubuntu 5.4.0-1045-raspi #49-Ubuntu SMP PREEMPT Wed Sep 29 17:49:16 UTC 2021 aarch64 aarch64 aarch64 GNU/Linux
```

``` sh
otool -t -v -V main > main-
```

[`panic.go deferproc`](https://github.com/golang/go/blob/go1.20.7/src/runtime/panic.go#L268-L299)

``` go
// Create a new deferred function fn, which has no arguments and results.
// The compiler turns a defer statement into a call to this.
func deferproc(fn func()) {
    // get g (goroutine のこと。)
	gp := getg()
    // m はマシンスレッドのこと。
	if gp.m.curg != gp {
		// go code on the system stack can't defer
		throw("defer on system stack")
	}

	// defer 本体の作成。
	d := newdefer()
	if d._panic != nil {
		throw("deferproc: d.panic != nil after newdefer")
	}
	// defer の next 呼び出しとして、今現在 g (goroutine) に乗ってる defer 関数を指定する。
	d.link = gp._defer
	// 今の goroutine の defer 関数として、今作成した defer 本体を再登録する。
	// つまり、最後に登録したものが最初に呼び出されることがわかる！
	gp._defer = d
	// 渡された関数を実行対象の関数として登録している。
	d.fn = fn
	d.pc = getcallerpc()
	// We must not be preempted between calling getcallersp and
	// storing it to d.sp because getcallersp's result is a
	// uintptr stack pointer.
	d.sp = getcallersp()

	// deferproc returns 0 normally.
	// a deferred func that stops a panic
	// makes the deferproc return 1.
	// the code the compiler generates always
	// checks the return value and jumps to the
	// end of the function if deferproc returns != 0.
	return0()
	// No code can go here - the C return register has
	// been set and must not be clobbered.
}
```

defer 本体は [runtime/runtime2.go](https://github.com/golang/go/blob/go1.20.7/src/runtime/runtime2.go#L981-L1013) にある

``` go
// A _defer holds an entry on the list of deferred calls.
// If you add a field here, add code to clear it in deferProcStack.
// This struct must match the code in cmd/compile/internal/ssagen/ssa.go:deferstruct
// and cmd/compile/internal/ssagen/ssa.go:(*state).call.
// Some defers will be allocated on the stack and some on the heap.
// All defers are logically part of the stack, so write barriers to
// initialize them are not required. All defers must be manually scanned,
// and for heap defers, marked.
type _defer struct {
	started bool
	heap    bool
	// openDefer indicates that this _defer is for a frame with open-coded
	// defers. We have only one defer record for the entire frame (which may
	// currently have 0, 1, or more defers active).
	openDefer bool
	sp        uintptr // sp at time of defer
	pc        uintptr // pc at time of defer
	fn        func()  // can be nil for open-coded defers
	_panic    *_panic // panic that is running defer
	link      *_defer // next defer on G; can point to either heap or stack!

	// If openDefer is true, the fields below record values about the stack
	// frame and associated function that has the open-coded defer(s). sp
	// above will be the sp for the frame, and pc will be address of the
	// deferreturn call in the function.
	fd   unsafe.Pointer // funcdata for the function associated with the frame
	varp uintptr        // value of varp for the stack frame
	// framepc is the current pc associated with the stack frame. Together,
	// with sp above (which is the sp associated with the stack frame),
	// framepc/sp can be used as pc/sp pair to continue a stack trace via
	// gentraceback().
	framepc uintptr
}
```

> Some defers will be allocated on the stack and some on the heap.

stack になることも heap になることもある。ほう。

defer を消費する部分は [runtime/panic.go deferreturn](https://github.com/golang/go/blob/go1.20.7/src/runtime/panic.go#L445-L478) にあります（linux でみた時と src でみた時でちょっと違う。）。

むちゃくちゃ短えです。

``` go
// deferreturn runs deferred functions for the caller's frame.
// The compiler inserts a call to this at the end of any
// function which calls defer.
func deferreturn() {
	gp := getg()
	for {
		d := gp._defer
		if d == nil {
			return
		}
		sp := getcallersp()
		if d.sp != sp {
			return
		}
		if d.openDefer {
			done := runOpenDeferFrame(d)
			if !done {
				throw("unfinished open-coded defers in deferreturn")
			}
			gp._defer = d.link
			freedefer(d)
			// If this frame uses open defers, then this
			// must be the only defer record for the
			// frame, so we can just return.
			return
		}

		fn := d.fn
		d.fn = nil
		gp._defer = d.link
		freedefer(d)
		fn()
	}
}
```

- [defer を解放する処理(freedefer)](https://github.com/golang/go/blob/go1.20.7/src/runtime/panic.go#L379-L431) もある
- [deferreturn の検索](https://github.com/search?q=repo%3Agolang%2Fgo+deferreturn&type=code&p=2)
- [G: goroutine 本体](https://github.com/golang/go/blob/go1.20.7/src/runtime/runtime2.go#L407-L506)
- [M: カーネルのマシンスレッド本体](https://github.com/golang/go/blob/go1.20.7/src/runtime/runtime2.go#L526-L607)
- [P: リソース本体](https://github.com/golang/go/blob/go1.20.7/src/runtime/runtime2.go#L609-L764)

## TODO

- return された時に担当するのはどこの誰？
