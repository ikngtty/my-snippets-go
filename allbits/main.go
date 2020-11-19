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
	for deci := 0; deci < 1<<bitsLen; deci++ {
		bits := make([]bool, bitsLen)
		for digit := 0; digit < bitsLen; digit++ {
			bits[bitsLen-1-digit] = deci&(1<<digit) > 0
		}

		callback(bits)
	}
}
