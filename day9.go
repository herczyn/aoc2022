package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data9")
	if err != nil {
		fmt.Println(err)
	}

	length := 10

	x := make([]int, length)
	y := make([]int, length)

	visited := map[int]map[int]bool{};

	visit := func(x int, y int) {
		if visited[x]==nil {
			visited[x]=map[int]bool{}
		}
		visited[x][y]=true
	}
	visit(0,0)

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();

		where := strings.Split(line, " ")[0]
		count, _ := strconv.Atoi(strings.Split(line, " ")[1])

		for i := 0; i < count; i++ {
			if (where == "R") {
				x[0]++;
			} else if (where == "L") {
				x[0]--;
			} else if (where == "U") {
				y[0]++;
			} else if (where == "D") {
				y[0]--;
			}
			for j:=1; j<length; j++ {
				if (x[j-1] == x[j]) {
					if y[j-1] == y[j]+2 {
						y[j] = y[j] + 1
					} else if y[j-1] == y[j]-2 {
						y[j] = y[j] - 1
					}
				} else if (y[j-1] == y[j]) {
					if x[j-1] == x[j]+2 {
						x[j] = x[j] + 1
					} else if x[j-1] == x[j]-2 {
						x[j] = x[j] - 1
					}
				} else if (x[j-1] == x[j] + 2) {
					if (y[j-1] > y[j]) {
						x[j] = x[j] + 1
						y[j] = y[j] + 1
					} else if (y[j-1] < y[j]) {
						x[j] = x[j] + 1
						y[j] = y[j] - 1
					}
				} else if (x[j-1] == x[j] - 2) {
					if (y[j-1] > y[j]) {
						x[j] = x[j] - 1
						y[j] = y[j] + 1
					} else if (y[j-1] < y[j]) {
						x[j] = x[j] - 1
						y[j] = y[j] - 1
					}
				} else if (y[j-1] == y[j] + 2) {
					if (x[j-1] > x[j]) {
						x[j] = x[j] + 1
						y[j] = y[j] + 1
					} else if (x[j-1] < x[j]) {
						x[j] = x[j] - 1
						y[j] = y[j] + 1
					}
				} else if (y[j-1] == y[j] - 2) {
					if (x[j-1] > x[j]) {
						x[j] = x[j] + 1
						y[j] = y[j] - 1
					} else if (x[j-1] < x[j]) {
						x[j] = x[j] - 1
						y[j] = y[j] - 1
					}
				}
			}
			visit(x[length-1], y[length-1])
		}
	}
	sum := 0
	for _, m := range visited {
		sum += len(m)
	}
	fmt.Println(sum)
}