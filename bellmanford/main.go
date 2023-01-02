package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func main() {
	const v = 4
	const e = 5
	edges := []Edge{{0, 1, 2}, {0, 2, 3}, {1, 2, -5}, {1, 3, 1}, {2, 3, 2}}

	minDists := MakeInts(v, MaxInt)
	minDists[0] = 0
	for i := 1; i <= v; i++ {
		updated := false
		for _, edge := range edges {
			if minDists[edge.From] == MaxInt {
				continue
			}
			dist := minDists[edge.From] + edge.Dist
			if dist < minDists[edge.To] {
				updated = true
				minDists[edge.To] = dist
			}
		}
		if !updated {
			break
		}
		if i == v {
			fmt.Println("NEGATIVE CYCLE")
			return
		}
	}

	for i := 0; i < v; i++ {
		if minDists[i] == MaxInt {
			fmt.Println("INF")
		} else {
			fmt.Println(minDists[i])
		}
	} // 0, 2, -3, -1
}

type Edge struct {
	From, To int
	Dist     int
}

// MakeInts returns a slice of the int array.
func MakeInts(length int, initVal int) []int {
	a := make([]int, length)

	if initVal != 0 {
		for i := 0; i < length; i++ {
			a[i] = initVal
		}
	}

	return a
}
