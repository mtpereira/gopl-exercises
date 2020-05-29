// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
// !+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// Loop returns the population count (number of set bits) of x using a loop.
func Loop(x uint64) int {
	var r byte
	for i := 0; i < 8; i++ {
		r += pc[byte(x>>(i*8))]
	}
	return int(r)
}

// Shift returns the population count (number of set bits) of x by shifting.
func Shift(x uint64) int {
	r := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			r++
		}
	}
	return r
}

// Clearing returns the population count (number of set bits) of x by clearing.
func Clearing(x uint64) int {
	r := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		r++
	}
	return r
}

// !-
