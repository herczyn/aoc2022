package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("data2")
	if err != nil {
		fmt.Println(err)
	}
	
	sum := 0;

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		// if line[2] == 'X' {
		// 	sum += 1;
		// }
		// if line[2] == 'Y' {
		// 	sum += 2;
		// }
		// if line[2] == 'Z' {
		// 	sum += 3;
		// }
		// if (line[0] == 'A' && line[2] == 'X') {
		// 	sum += 3;
		// }
		// if (line[0] == 'B' && line[2] == 'Y') {
		// 	sum += 3;
		// }
		// if (line[0] == 'C' && line[2] == 'Z') {
		// 	sum += 3;
		// }
		// if (line[0] == 'A' && line[2] == 'Y') {
		// 	sum += 6;
		// }
		// if (line[0] == 'B' && line[2] == 'Z') {
		// 	sum += 6;
		// }
		// if (line[0] == 'C' && line[2] == 'X') {
		// 	sum += 6;
		// }
		var val int;
		if (line[0] == 'A') {
			val = 1;
		}
		if (line[0] == 'B') {
			val = 2;
		}
		if (line[0] == 'C') {
			val = 3;
		}
		if line[2] == 'X' {
			sum += 0;
			val += 2;
		}
		if line[2] == 'Y' {
			sum += 3;
			val += 0;
		}
		if line[2] == 'Z' {
			sum += 6;
			val += 1;
		}
		if (val > 3) {
			val -= 3;
		}
		sum += val;
	}
	fmt.Println(sum);
}




