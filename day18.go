package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data18")
	if err != nil {
		fmt.Println(err)
	}

	cubes := map[int]map[int]map[int]bool{}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	unset:=true
	var minX, maxX, minY, maxY, minZ, maxZ int
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		x, _:= strconv.Atoi(strings.Split(line, ",")[0])
		y, _:= strconv.Atoi(strings.Split(line, ",")[1])
		z, _:= strconv.Atoi(strings.Split(line, ",")[2])
		fmt.Println(x, y, z)
		if (cubes[x] == nil) {
			cubes[x] = map[int]map[int]bool{}
		}
		if (cubes[x][y] == nil) {
			cubes[x][y] = map[int]bool{}
		}
		cubes[x][y][z] = true
		if (unset) {
			minX = x
			maxX = x
			minY = y
			maxY = y
			minZ = z
			maxZ = z
			unset = false
		} else {
			if (x < minX) {
				minX = x
			}
			if (x > maxX) {
				maxX = x
			}
			if (y < minY) {
				minY = y
			}
			if (y > maxY) {
				maxY = y
			}
			if (z < minZ) {
				minZ = z
			}
			if (z > maxZ) {
				maxZ = z
			}
		}
	}

	fmt.Println(minX, maxX, minY, maxY, minZ, maxZ)

	sides := 0

	outside := map[int]map[int]map[int]int{}

	var isOutside func(x int, y int, z int) (bool);
	isOutside = func(x int, y int, z int) (bool) {
		if cubes[x][y][z] {
			return false
		}
		// return true
		if x<=minX || x>=maxX || y<=minY || y>=maxY || z<=minZ || z>=maxZ {
			return true
		}
		if (outside[x][y][z]==0) {
			if (outside[x] == nil) {
				outside[x] = map[int]map[int]int{}
			}
			if (outside[x][y] == nil) {
				outside[x][y] = map[int]int{}
			}
		    outside[x][y][z] = -1
			if isOutside(x+1, y, z) || isOutside(x-1, y, z) ||
			   isOutside(x, y+1, z) || isOutside(x, y-1, z) ||
			   isOutside(x, y, z+1) || isOutside(x, y, z-1) {
		    	outside[x][y][z] = 1
		    }
		}
		return outside[x][y][z] == 1
	}

	for x, _ := range cubes {
		for y, _ := range cubes[x] {
			for z, _ := range cubes[x][y] {
				if (isOutside(x+1, y, z)) {
					sides++
				}
				if (isOutside(x-1, y, z)) {
					sides++
				}
				if (isOutside(x, y+1, z)) {
					sides++
				}
				if (isOutside(x, y-1, z)) {
					sides++
				}
				if (isOutside(x, y, z+1)) {
					sides++
				}
				if (isOutside(x, y, z-1)) {
					sides++
				}
				fmt.Println(x, y, z, sides)
			}
		}
	}

	fmt.Println(sides)

}
