package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
)

func main() {
	readFile, err := os.Open("data1")
    if err != nil {
        fmt.Println(err)
    }
	
	subsum := 0;
	var allsums []int;

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
    	line := fileScanner.Text();
		if (line == "") {
			allsums = append(allsums, subsum);
			subsum = 0;
		} else {
			val, err := strconv.Atoi(line);
		    if err != nil {
		        fmt.Println(err)
		    }
		    subsum += val;
		}
	}
	allsums = append(allsums, subsum);
	sort.Ints(allsums)
	fmt.Println(allsums);
}




