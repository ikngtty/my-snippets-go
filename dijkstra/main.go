package main

import (
	"container/heap"
	"fmt"
)

func main() {
	const v = 4
	const e = 5
	g := make([][]Edge, v)
	g[0] = []Edge{{0, 1, 1}, {0, 2, 4}}
	g[1] = []Edge{{1, 2, 2}, {1, 3, 5}}
	g[2] = []Edge{{2, 3, 1}}

	minDists := MakeInts(v, -1)
	hv := HeapVertex{Vertex{0, 0}}
	for len(hv) > 0 {
		v := heap.Pop(&hv).(Vertex)
		if minDists[v.Node] >= 0 {
			continue
		}
		minDists[v.Node] = v.Dist

		for _, edge := range g[v.Node] {
			heap.Push(&hv, Vertex{edge.To, v.Dist + edge.Dist})
		}
	}

	for i := 0; i < v; i++ {
		fmt.Println(minDists[i]) // 0, 1, 3, 4
	}
}

type Edge struct {
	From, To int
	Dist     int
}

type Vertex struct {
	Node int
	Dist int
}

type HeapVertex []Vertex

func (h HeapVertex) Len() int           { return len(h) }
func (h HeapVertex) Less(i, j int) bool { return h[i].Dist < h[j].Dist }
func (h HeapVertex) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *HeapVertex) Push(x interface{}) {
	*h = append(*h, x.(Vertex))
}

func (h *HeapVertex) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
