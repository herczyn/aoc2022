package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("data17")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	createObject := func(i int, top int) [][]int {
		top = top+4
		if (i%5==0) {
			return [][]int{[]int{2, top}, []int{3, top}, []int{4, top}, []int{5, top}}
		}
		if (i%5==1) {
			return [][]int{[]int{2, top+1}, []int{3, top}, []int{3, top+1}, []int{3, top+2}, []int{4, top+1}}
		}
		if (i%5==2) {
			return [][]int{[]int{2, top}, []int{3, top}, []int{4, top}, []int{4, top+1}, []int{4, top+2}}
		}
		if (i%5==3) {
			return [][]int{[]int{2, top}, []int{2, top+1}, []int{2, top+2}, []int{2, top+3}}
		}
		if (i%5==4) {
			return [][]int{[]int{2, top}, []int{2, top+1}, []int{3, top}, []int{3, top+1}}
		}
		return nil
	}

	fileScanner.Scan()
	winds := fileScanner.Text();

	height := -1
	var object [][]int
	pos := 0
	objectNo := 0
	occupied := map[int]map[int]bool{}

	way := func() (int) {
		if winds[pos % len(winds)] == '>' {
			return +1
		} else {
			return -1
		}
	}

	lastHeight := 0
	lastObject := 0

	move := func() {
		if pos % len(winds) == 0 && pos>0 {
			fmt.Println("height: ", height, height - lastHeight)
			lastHeight = height
			fmt.Println("object: ", objectNo, objectNo - lastObject)
			lastObject = objectNo
		}
		if (objectNo - lastObject) == 1188 {
			fmt.Println(height - lastHeight)
		}
		if object==nil {
			object = createObject(objectNo, height)
		}
		canMove := true
		for _, point := range object {
			if occupied[point[0]+way()][point[1]] {
				canMove = false
			}
			if point[0]+way()<0 || point[0]+way()>6 {
				canMove = false
			}
		}
		if canMove {
			for i, point := range object {
				object[i] = []int{point[0]+way(), point[1]}
			}
		}
		pos++
		canMove = true
		for _, point := range object {
			if occupied[point[0]][point[1]-1] {
				canMove = false
			}
			if point[1]-1<0 {
				canMove = false
			}
		}
		if canMove {
			for i, point := range object {
				object[i] = []int{point[0], point[1]-1}
			}
		} else {
			for _, point := range object {
				if (occupied[point[0]] == nil) {
					occupied[point[0]] = map[int]bool{}
				}
				occupied[point[0]][point[1]] = true
				if height<point[1] {
					height = point[1]
				}
			}
			object = nil
			objectNo++
		}
	}

	for objectNo < 5000 {
		move()
	}

	fmt.Println(height+1)

	height = 2642
	objectNo = 1732

	fmt.Println((1000000000000 - 1732)%1740)
	fmt.Println(2642+((1000000000000 - 1732 - 1188)/1740)*2666+1833+1)
}

