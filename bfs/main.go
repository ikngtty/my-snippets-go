package main

import "fmt"

func main() {
	const h = 5
	const w = 6
	start := Point{0, 0}
	goal := Point{h - 1, w - 1}

	maze := make([]string, h)
	maze[0] = "...##."
	maze[1] = ".###.."
	maze[2] = ".##..."
	maze[3] = "....#."
	maze[4] = "....#."

	dy := []int{1, 0, -1, 0}
	dx := []int{0, 1, 0, -1}

	visitedAt := Make2DInts(h, w, -1)
	visitedAt[start.Y][start.X] = 0
	visitedQueue := NewListPoint()
	visitedQueue.Push(start)
	for visitedQueue.Len() > 0 {
		cur := visitedQueue.PopLeft()
		for i := 0; i < 4; i++ {
			next := Point{cur.Y + dy[i], cur.X + dx[i]}
			if next.Y >= 0 && next.Y < h && next.X >= 0 && next.X < w &&
				maze[next.Y][next.X] == '.' &&
				visitedAt[next.Y][next.X] < 0 {
				visitedAt[next.Y][next.X] = visitedAt[cur.Y][cur.X] + 1
				visitedQueue.Push(next)
			}
		}
	}

	fmt.Println(visitedAt[goal.Y][goal.X]) // 11
}

type Point struct {
	Y, X int
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

// ListPoint is a list of Point.
type ListPoint struct {
	first *listPointNode
	last  *listPointNode
	len   int
}

type listPointNode struct {
	parent *listPointNode
	child  *listPointNode
	value  Point
}

// NewListPoint returns a new ListPoint.
func NewListPoint() *ListPoint {
	return &ListPoint{nil, nil, 0}
}

// NewListPointFromArray returns a new ListPoint having the values `a` has.
func NewListPointFromArray(a []Point) *ListPoint {
	list := NewListPoint()
	for _, elem := range a {
		list.Push(elem)
	}
	return list
}

// Len returns the length of the list.
func (list *ListPoint) Len() int {
	return list.len
}

// Push pushes elem to the end of the list.
func (list *ListPoint) Push(elem Point) {
	node := listPointNode{list.last, nil, elem}
	if list.first == nil {
		list.first = &node
	} else {
		list.last.child = &node
	}
	list.last = &node
	list.len++
}

// PushLeft pushes elem to the beginning of the list.
func (list *ListPoint) PushLeft(elem Point) {
	node := listPointNode{nil, list.first, elem}
	if list.last == nil {
		list.last = &node
	} else {
		list.first.parent = &node
	}
	list.first = &node
	list.len++
}

// Pop pops elem from the end of the list.
func (list *ListPoint) Pop() Point {
	if list.last == nil {
		panic("no item")
	}
	value := list.last.value
	list.last = list.last.parent
	if list.last == nil {
		list.first = nil
	} else {
		list.last.child = nil
	}
	list.len--
	return value
}

// PopLeft pops elem from the beginning of the list.
func (list *ListPoint) PopLeft() Point {
	if list.first == nil {
		panic("no item")
	}
	value := list.first.value
	list.first = list.first.child
	if list.first == nil {
		list.last = nil
	} else {
		list.first.parent = nil
	}
	list.len--
	return value
}

// Concat concatenates the list and the other.
func (list *ListPoint) Concat(other *ListPoint) {
	if list.first == nil {
		*list = *other
	} else if other.first == nil {
		*other = *list
	} else {
		list.last.child = other.first
		other.first.parent = list.last
		list.last = other.last
		other.first = list.first
		list.len += other.len
		other.len = list.len
	}
}

// Each applies f for every element in the list.
func (list *ListPoint) Each(f func(elem Point)) {
	cur := list.first
	for cur != nil {
		f(cur.value)
		cur = cur.child
	}
}

// ToA converts the list to an array.
func (list *ListPoint) ToA() []Point {
	a := make([]Point, list.len)
	{
		index := 0
		list.Each(func(elem Point) {
			a[index] = elem
			index++
		})
	}
	return a
}
