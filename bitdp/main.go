package main

import "fmt"

func main() {
	const v = 4
	g := Make2DInts(v, v, -1)
	g[0][1] = 2
	g[1][2] = 3
	g[1][3] = 9
	g[2][0] = 1
	g[2][3] = 6
	g[3][2] = 4

	table := Make2DInts(1<<uint(v), v, -1)
	table[0][0] = 0
	for s := 0; s < 1<<uint(v); s++ {
		for from := 0; from < v; from++ {
			if table[s][from] < 0 {
				continue
			}

			for to := 0; to < v; to++ {
				if s&(1<<uint(to)) > 0 || g[from][to] < 0 {
					continue
				}

				d := table[s][from] + g[from][to]
				toS := s | (1 << uint(to))
				if table[toS][to] < 0 || d < table[toS][to] {
					table[toS][to] = d
				}
			}
		}
	}

	fmt.Println(table[1<<uint(v)-1][0]) // 16
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
