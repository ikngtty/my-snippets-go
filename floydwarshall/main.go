package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)

func main() {
	const v = 4
	const e = 6
	edges := []Edge{{0, 1, 1}, {0, 2, 5}, {1, 2, 2}, {1, 3, 4}, {2, 3, 1}, {3, 2, 7}}

	table := Make2DInts(v, v, MaxInt)
	for i := 0; i < v; i++ {
		table[i][i] = 0
	}
	for _, edge := range edges {
		table[edge.From][edge.To] = edge.Dist
	}
	for k := 0; k < v; k++ {
		for i := 0; i < v; i++ {
			for j := 0; j < v; j++ {
				if table[i][k] < MaxInt && table[k][j] < MaxInt {
					table[i][j] = Min(table[i][j], table[i][k]+table[k][j])
				}
			}
		}
	}

	for i := 0; i < v; i++ {
		if table[i][i] < 0 {
			fmt.Println("NEGATIVE CYCLE")
			return
		}
	}
	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			if table[i][j] == MaxInt {
				fmt.Print("INF")
			} else {
				fmt.Print(table[i][j])
			}
		}
		fmt.Println()
	}
	// 0 1 3 4
	// INF 0 2 3
	// INF INF 0 1
	// INF INF 7 0
}

type Edge struct {
	From, To int
	Dist     int
}

// Min returns the minimum value of the specified values.
func Min(values ...int) int {
	if len(values) == 0 {
		panic("no values")
	}

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

// Make2DInts returns a slice of the two-dimensional int array.
func Make2DInts(xLen, yLen int, initVal int) [][]int {
	a := make([][]int, xLen)
	for x := 0; x < xLen; x++ {
		a[x] = make([]int, yLen)
	}

	if initVal != 0 {
		for x := 0; x < xLen; x++ {
			for y := 0; y < yLen; y++ {
				a[x][y] = initVal
			}
		}
	}

	return a
}
