package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data14")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	pos := map[int]map[int]bool{}
	abyss := 0

	for fileScanner.Scan() {
		line := fileScanner.Text();
		points := strings.Split(line, " -> ")
		for i:=0; i<len(points)-1; i++ {
			point1 := strings.Split(points[i],   ",")
			point2 := strings.Split(points[i+1], ",")

			x1, _ := strconv.Atoi(point1[0])
			y1, _ := strconv.Atoi(point1[1])
			x2, _ := strconv.Atoi(point2[0])
			y2, _ := strconv.Atoi(point2[1])

			if y1>abyss {
				abyss = y1
			}
			if y2>abyss {
				abyss = y2
			}

			for x1!=x2 || y1!=y2 {
				if (pos[x1]==nil) {
					pos[x1] = map[int]bool{}
				}
				pos[x1][y1] = true
				if x1>x2 {
					x1-=1 
				}
				if x1<x2 {
					x1+=1 
				}
				if y1>y2 {
					y1-=1 
				}
				if y1<y2 {
					y1+=1 
				}
			}
			if (pos[x2]==nil) {
				pos[x2] = map[int]bool{}
			}
			pos[x2][y2] = true
		}
	}
	// count:=0
	// for {
	// 	x:=500
	// 	y:=0
	// 	for y<abyss {
	// 		if !pos[x][y+1] {
	// 			y++
	// 		} else if !pos[x-1][y+1] {
	// 			x--
	// 			y++
	// 		} else if !pos[x+1][y+1] {
	// 			x++
	// 			y++
	// 		} else {
	// 			break
	// 		}
	// 	}
	// 	if y==abyss {
	// 		break
	// 	}
	// 	if (pos[x]==nil) {
	// 		pos[x] = map[int]bool{}
	// 	}
	// 	pos[x][y] = true
	// 	count++
	// }
	// fmt.Println(count)
	count:=0
	for {
		x:=500
		y:=0
		for y<abyss+1 {
			if !pos[x][y+1] {
				y++
			} else if !pos[x-1][y+1] {
				x--
				y++
			} else if !pos[x+1][y+1] {
				x++
				y++
			} else {
				break
			}
		}
		if (y==0) {
			count++
			break
		}
		if (pos[x]==nil) {
			pos[x] = map[int]bool{}
		}
		pos[x][y] = true
		count++
	}
	fmt.Println(count)
}
