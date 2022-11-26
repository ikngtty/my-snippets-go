package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50, 60, 70}

	ok := len(a)
	ng := -1
	for ok-ng > 1 {
		mid := (ok + ng) / 2
		if a[mid] > 42 {
			ok = mid
		} else {
			ng = mid
		}
	}

	fmt.Println(a[ok])
}
