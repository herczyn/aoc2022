package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data10")
	if err != nil {
		fmt.Println(err)
	}

	val := 1
	sum := 0
	cyc := 0
	rows := []string{}

	cycle := func() {
		cyc++
		fmt.Println(cyc, cyc%40, val)
		if cyc % 40 == 20 {
			sum += val * cyc
		}
		if cyc % 40 == 1 {
			rows = append(rows, "")
		}
		if (cyc % 40 >= val - 1 && cyc % 40 <= val + 1) {
			rows[len(rows) - 1] += "#"
		} else {
			rows[len(rows) - 1] += "."
		}

	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();

		what := strings.Split(line, " ")[0]
		diff := 0
		if what != "noop" {
			diff, _ = strconv.Atoi(strings.Split(line, " ")[1])
		}

		if what == "noop" {
			cycle()
		}
		if what == "addx" {
			cycle()
			val+=diff
			cycle()
			//val+=diff
		}
	}
	fmt.Println(sum)
	for _, r := range rows {
		fmt.Println(r)
	}
}