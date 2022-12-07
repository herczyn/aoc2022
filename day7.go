package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
    "log"
)

type Dir struct {
	dirs map[string]Dir
	files map[string]int
}

func main() {
	readFile, err := os.Open("data7")
	if err != nil {
		fmt.Println(err)
	}
	
	root := Dir{map[string]Dir{}, map[string]int{}}
	stack := []Dir{root}
	inls := false

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();

		if (inls && !strings.HasPrefix(line, "$")) {
			if (strings.HasPrefix(line, "dir")) {
				stack[0].dirs[line[4:]] = Dir{map[string]Dir{}, map[string]int{}}
			} else {
				split := strings.Split(line, " ")
				stack[0].files[split[1]], _ = strconv.Atoi(split[0])
			}
		} else {
			inls = false
			if (line == "$ cd /") {
				stack = []Dir{root}
			} else if (line == "$ cd ..") {
				stack = stack[1:]
			} else if (strings.HasPrefix(line, "$ cd ")) {
				stack = append([]Dir{stack[0].dirs[line[5:]]}, stack...)
			} else if (line == "$ ls") {
				inls = true
			} else {
        		log.Fatal("unknown line: " + line)
			}
		}
		fmt.Println(root)
	}
	sum:=0
	newSizes, totalSize := sizes(root)
	fmt.Println(newSizes)
	fmt.Println(totalSize)
	for _,size := range newSizes {
		if size <= 100000 {
			sum += size
		}
		if (size >= 6552309) {
			fmt.Println("option:", size)
		}
	}
	fmt.Println(sum)
}

func sizes(dir Dir) ([]int, int) {
	newSizes := []int{}
	totalSize := 0
	for _, d := range dir.dirs {
		sizes, size := sizes(d)
		newSizes = append(newSizes, sizes...)
		totalSize += size
	}
	for _, s := range dir.files {
		totalSize += s
	}
	newSizes = append(newSizes, totalSize)
	return newSizes, totalSize
}


