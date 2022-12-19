package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data11")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	items := map[int][]int{}
	operation := map[int][]string{}
	divisible := map[int]int{}
	whentrue := map[int]int{}
	whenfalse := map[int]int{}
	inspections := map[int]int{}
  
	for fileScanner.Scan() {
		line1 := fileScanner.Text();
		fileScanner.Scan()
		line2 := fileScanner.Text();
		fileScanner.Scan()
		line3 := fileScanner.Text();
		fileScanner.Scan()
		line4 := fileScanner.Text();
		fileScanner.Scan()
		line5 := fileScanner.Text();
		fileScanner.Scan()
		line6 := fileScanner.Text();
		fileScanner.Scan()

		no, _ := strconv.Atoi(strings.Split(strings.Split(line1, " ")[1], ":")[0])
		string := strings.Split(strings.Split(line2, ": ")[1], ", ")
		item := make([]int, len(string))
		for i, s := range string {
			item[i], _ = strconv.Atoi(s)
		}
		items[no] = item
		operation[no] = strings.Split(strings.Split(line3, "= old ")[1], " ")
		divisible[no], _ = strconv.Atoi(strings.Split(line4, "divisible by ")[1])
		whentrue[no], _ = strconv.Atoi(strings.Split(line5, "to monkey ")[1])
		whenfalse[no], _ = strconv.Atoi(strings.Split(line6, "to monkey ")[1])
	}
	d := 1
	for _, div := range divisible {
		d*=div
	}
	fmt.Println(d)
	for r := 0; r < 10000; r++ {
		for no :=0; no < len(operation); no++ {
			for _, worry := range items[no] {
				inspections[no] += 1;
				if operation[no][0]== "+" {
					val, _ := strconv.Atoi(operation[no][1])
					worry += val
				} else if operation[no][1]== "old" {
					worry *= worry
				} else {
					val, _ := strconv.Atoi(operation[no][1])
					worry *= val
				}
				// worry/=3
				worry%=d
				if (worry % divisible[no] == 0) {
					items[whentrue[no]] = append(items[whentrue[no]], worry)
				} else {
					items[whenfalse[no]] = append(items[whenfalse[no]], worry)
				}
			}
			items[no] = []int{}
		}

		fmt.Println(items)
		fmt.Println(inspections)
	}
}





