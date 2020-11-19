package main

import "fmt"

func main() {
	btoi := map[bool]int{false: 0, true: 1}

	AllBits(5, func(bits []bool) {
		for _, b := range bits {
			fmt.Printf("%d", btoi[b])
		}
		fmt.Println()
	})
}

func AllBits(bitsLen int, callback func(bits []bool)) {
	for pattern := 0; pattern < 1<<bitsLen; pattern++ {
		bits := make([]bool, bitsLen)
		for rpos := 0; rpos < bitsLen; rpos++ {
			bits[bitsLen-1-rpos] = pattern&(1<<rpos) > 0
		}

		callback(bits)
	}
}
