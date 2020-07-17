package main

import (
	"fmt"
)

const (
	IB = 1
	KB = IB * 1000
	MB = KB * 1000
	GB = MB * 1000
	TB = GB * 1000
	PB = TB * 1000
	EB = PB * 1000
	ZB = EB * 1000
	YB = ZB * 1000.0
)

func main() {
	fmt.Println(KB)
	fmt.Println(YB)
}
