package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data4")
	if err != nil {
		fmt.Println(err)
	}
	
	sum := 0;

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		l1 := line[:strings.Index(line, ",")]
		l2 := line[(strings.Index(line, ",")+1):]
		s1, _ := strconv.Atoi(l1[:strings.Index(l1, "-")])
		e1, _ := strconv.Atoi(l1[(strings.Index(l1, "-")+1):])
		s2, _ := strconv.Atoi(l2[:strings.Index(l2, "-")])
		e2, _ := strconv.Atoi(l2[(strings.Index(l2, "-")+1):])
		
		// if ((s1>=s2 && e1<=e2) || (s2>=s1 && e2<=e1)) {
		if ((s1<=e2 && e1>=s2) || (s2<=e1 && e2>=s1)) {
			sum+=1;
		}
		fmt.Println(s1, e1, s2, e2);
	}
	fmt.Println(sum);
}




