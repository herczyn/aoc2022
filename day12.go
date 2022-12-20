package main

import (
	"bufio"
	"fmt"
	"os"
	"container/heap"
	"strings"
)

func main() {
	readFile, err := os.Open("data12")
	if err != nil {
		fmt.Println(err)
	}

	m := []string{}
	h := &IntHeap{}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)


	x := 0
	y := 0
  
  	r := 0
	for fileScanner.Scan() {
		line := fileScanner.Text();
		m = append(m, line)
		if strings.Contains(line, "E") {
			x = r
			y = len(strings.Split(line, "E")[0])
		}
		r++
	}

	fmt.Println("starting", x, y)

	heap.Push(h, 1)
	p := map[int][][]int{1: [][]int{[]int{x, y}}}

	visited := map[int]map[int]bool{}

	comp := func(c byte, d byte) (bool) {
		if c == 'S' {
			c = 'a'
		}
		if c == 'E' {
			c = 'z'
		}
		if d == 'S' {
			d = 'a'
		}
		if d == 'E' {
			d = 'z'
		}
		return c >= d - 1
	}

	move := func() (bool) {
		distance := heap.Pop(h).(int)
		position := p[distance][0]
		p[distance] = p[distance][1:]

		x = position[0]
		y = position[1]
		if (m[x][y] == 'S' || m[x][y] == 'a') {
			fmt.Println("FOUND IT ", distance - 1)
			return true
		}
		if (visited[x][y]) {
			return false
		}
		if (visited[x] == nil) {
			visited[x] = map[int]bool{}
		}
		visited[x][y] = true
		if (p[distance + 1] == nil) {
			p[distance + 1] = [][]int{}
		}
		if (x>0 && comp(m[x-1][y], m[x][y]) && !visited[x-1][y]) {
			heap.Push(h, distance + 1)
			p[distance + 1] = append(p[distance + 1], []int{x-1, y})
		}
		if (x<len(m)-1 && comp(m[x+1][y], m[x][y]) && !visited[x+1][y]) {
			heap.Push(h, distance + 1)
			p[distance + 1] = append(p[distance + 1], []int{x+1, y})
		}
		if (y>0 && comp(m[x][y-1], m[x][y]) && !visited[x][y-1]) {
			heap.Push(h, distance + 1)
			p[distance + 1] = append(p[distance + 1], []int{x, y-1})
		}
		if (y<len(m[x])-1 && comp(m[x][y+1], m[x][y]) && !visited[x][y+1]) {
			heap.Push(h, distance + 1)
			p[distance + 1] = append(p[distance + 1], []int{x, y+1})
		}
		return false
	}

	for !move() {
	}

}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

