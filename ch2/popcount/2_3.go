package main

import (
	"fmt"
	"os"
	"strconv"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountLoop is PopCount but uses a loop.
func PopCountLoop(x uint64) int {
	var b byte
	var i uint64
	for i = 0; i < 8; i++ {
		b += pc[byte(x>>(i*8))]
	}
	return int(b)
}

// ShiftCount is PopCount but uses a bit shift.
func ShiftCount(x uint64) int {
	var ans uint64
	var i uint64
	for i = 0; i < 64; i++ {
		x >>= 1
		ans += x & 1
	}
	return int(ans)
}

// QuickCount is PopCount but pops the right most 1 bit.
func QuickCount(x uint64) int {
	var i int
	for i = 0; i < 64; i++ {
		if x == 0 {
			return i
		}
		x = x & (x - 1)
	}
	return i
}

func main() {
	i, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err == nil {
		ans := QuickCount(i)
		fmt.Printf("%d\n", ans)
	}
}
