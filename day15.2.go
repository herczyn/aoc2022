package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data15")
	if err != nil {
		fmt.Println(err)
	}

	max := 4000000

	taken := make(map[int]map[int]int)

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	set := func(x int, y1 int, y2 int) {
		x = Max(Min(x, max), 0)
		y1 = Max(Min(y1, max), 0)
		y2 = Max(Min(y2, max), 0)
		if taken[x]==nil {
			taken[x] = make(map[int]int)
		}
		for k, v := range taken[x] {
			if (k != y1 || v != y2) && ((k>=y1-1 && k<=y2+1) || (y1>=k-1 && y1<=v+1)) {
				delete(taken[x], k)
				taken[x][Min(k, y1)] = Max(v, y2)
				return
			}
		}
		taken[x][y1] = y2
	}
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		split := strings.Split(line, "=")
		sx, _ := strconv.Atoi(strings.Split(split[1], ",")[0])
		sy, _ := strconv.Atoi(strings.Split(split[2], ":")[0])
		bx, _ := strconv.Atoi(strings.Split(split[3], ",")[0])
		by, _ := strconv.Atoi(split[4])
		distance := Abs(sx-bx) + Abs(sy-by)
		fmt.Println(sx, sy, bx, by, distance)
		for i := -distance; i <= distance; i++ {
			set(sx+i, sy-Abs(Abs(i)-distance), sy+Abs(Abs(i)-distance))
		}
	}
	for i:=0; i<=max; i++ {
		for (len(taken[i]) > 2) {
			for k, v := range(taken[i]) {
				set(i, k, v)
			}
		}
		for k, v := range(taken[i]) {
			set(i, k, v)
		}
		for k, v := range(taken[i]) {
			set(i, k, v)
		}
		if (len(taken[i]) > 1) {
			fmt.Println(i, taken[i])
		}
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}


func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
