package main

import (
	"testing"
)

func Benchmark_run(b *testing.B) {
	for n := 0; n < b.N; n++ {
		run()
	}
}
