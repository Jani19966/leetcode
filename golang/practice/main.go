package main

import (
	"fmt"
	"os"
	"sync/atomic"
)

func testAtomicReference(input *atomic.Bool) {
	input.Store(true)
	fmt.Println(input.Load())
}

type test struct {
	name             string
	pointerName      *string
	age              int
	pointerAge       *int
	shouldRun        atomic.Bool
	shouldRunPointer *atomic.Bool
}

// Test how the == and similars work
func testTypeEquals() {

}

func main() {
	var input atomic.Bool
	input.Store(false)
	testAtomicReference(&input)

	var num *int
	x := 10
	num = &x
	*num = 22

	fmt.Println(num, &x)

	fmt.Fprintf(os.Stdout, "Test struct: %v \n", []any{
		test{
			name:             "",
			pointerName:      new(string),
			age:              0,
			pointerAge:       new(int),
			shouldRun:        atomic.Bool{},
			shouldRunPointer: &atomic.Bool{},
		},
	}...)

}
