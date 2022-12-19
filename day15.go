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

	row := 2000000

	taken := make(map[int]bool)
	beacons := make(map[int]bool)

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		split := strings.Split(line, "=")
		sx, _ := strconv.Atoi(strings.Split(split[1], ",")[0])
		sy, _ := strconv.Atoi(strings.Split(split[2], ":")[0])
		bx, _ := strconv.Atoi(strings.Split(split[3], ",")[0])
		by, _ := strconv.Atoi(split[4])
		if (by == row) {
			beacons[bx] = true
		}
		distance := Abs(sx-bx) + Abs(sy-by) - Abs(sy-row)
		if (distance > 0) {
			taken[sx] = true
			for i := 1; i <= distance; i++ {
				taken[sx+i] = true
				taken[sx-i] = true
			}
		}

		fmt.Println(sx, sy, bx, by)
	}
	fmt.Println(len(taken) - len(beacons))
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
