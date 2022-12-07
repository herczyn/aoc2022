package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func main() {
	L := 14
	readFile, err := os.Open("data6")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();

		for i := L-1; i<len(line); i++ {
			found := false
			for j := 0; j < L-1; j++ {
				for k := j+1; k < L; k++ {
					if line[i-j] == line[i-k] {
						found = true
					}
				}
			}
			if (!found) {
				fmt.Println(i+1)
				return
			}
		}
	}
}

