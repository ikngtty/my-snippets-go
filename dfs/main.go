package main

import "fmt"

func main() {
	const h = 5
	const w = 6

	atlas := make([][]byte, h)
	atlas[0] = []byte("111011")
	atlas[1] = []byte("111011")
	atlas[2] = []byte("100000")
	atlas[3] = []byte("101001")
	atlas[4] = []byte("101111")

	dy := []int{1, 0, -1, 0}
	dx := []int{0, 1, 0, -1}

	var dfs func(cur Point)
	dfs = func(cur Point) {
		atlas[cur.Y][cur.X] = '0'
		for i := 0; i < 4; i++ {
			next := Point{cur.Y + dy[i], cur.X + dx[i]}
			if next.Y >= 0 && next.Y < h && next.X >= 0 && next.X < w &&
				atlas[next.Y][next.X] == '1' {
				dfs(next)
			}
		}
	}

	count := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if atlas[y][x] == '1' {
				dfs(Point{y, x})
				count++
			}
		}
	}

	fmt.Println(count) // the connected component count is 3
}

type Point struct {
	Y, X int
}
