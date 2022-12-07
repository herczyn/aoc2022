package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data5")
	if err != nil {
		fmt.Println(err)
	}
	
	moves := false;
	crates := make([][]string, 10)

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();

		if (line == "") {
			moves = true;
		} else if (moves) {
			split := strings.Split(line, " ")
			fmt.Println(split[1], split[3], split[5])
			n, _ := strconv.Atoi(split[1])
			from, _ := strconv.Atoi(split[3])
			to, _ := strconv.Atoi(split[5])
			// for i:=0; i<n; i++ {
				crates[to-1] = append(append([]string{}, crates[from-1][:n]...), crates[to-1]...)
				crates[from-1] = crates[from-1][n:]
			// }
			fmt.Println(crates)
		} else if (line[1]!='1') {
			for i:=0; i<9; i++ {
				if (line[i*4+1]!= ' ') {
					crates[i] = append(crates[i], line[i*4+1:i*4+2])
				}
			}
			fmt.Println(crates)
		}
	}
}




