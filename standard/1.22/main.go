package main

import (
	"cmp"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	host := cmp.Or(os.Getenv("HOST"), "localhost:8080")
	fmt.Printf("host: %v\n", host)

	os.Setenv("HOST", "localhost:15555")
	host2 := cmp.Or(os.Getenv("HOST"), "localhost:8080")
	fmt.Printf("host: %v\n", host2)

	nonzero := cmp.Or(0, 0, 134)
	fmt.Printf("nonzero: %v\n", nonzero)

	start := time.Now()

	defer fmt.Printf("time.Since(start): %v\n", time.Since(start))

	defer func() {
		fmt.Printf("time.Since(start): %v\n", time.Since(start))
	}()

	time.Sleep(500 * time.Millisecond)

	x := 1
	defer fmt.Printf("deferred x: %v\n", x)
	defer func(x *int) {
		fmt.Printf("deferred xp: %v\n", *x)
	}(&x)

	x++
	fmt.Printf("x: %v\n", x)
}

func deferr() {
	runtime.Gosched()
}

// runtime/panic.go
//
// // Create a new deferred function fn, which has no arguments and results.
// // The compiler turns a defer statement into a call to this.
// func deferproc(fn func()) {
// 	gp := getg()
// 	if gp.m.curg != gp {
// 		// go code on the system stack can't defer
// 		throw("defer on system stack")
// 	}
// 	d := newdefer()
// 	d.link = gp._defer
// 	gp._defer = d
// 	d.fn = fn
// 	d.pc = getcallerpc()
// 	// We must not be preempted between calling getcallersp and
// 	// storing it to d.sp because getcallersp's result is a
// 	// uintptr stack pointer.
// 	d.sp = getcallersp()
// 	// deferproc returns 0 normally.
// 	// a deferred func that stops a panic
// 	// makes the deferproc return 1.
// 	// the code the compiler generates always
// 	// checks the return value and jumps to the
// 	// end of the function if deferproc returns != 0.
// 	return0()
// 	// No code can go here - the C return register has
// 	// been set and must not be clobbered.
// }
