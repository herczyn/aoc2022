package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
	readFile, err := os.Open("data3")
    if err != nil {
        fmt.Println(err)
    }
	
	sum := 0;

    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
    	// line := fileScanner.Text();
    	// for i := 0; i < len(line) / 2; i++ {
    	// 	index := strings.Index(line[(len(line) / 2):], line[i:i+1])
    	// 	if (index >= 0) {
    	// 		if (line[i] >= 'a') {
		//     		sum += int(line[i] - 'a') + 1
		//     	} else {
		//     		sum += int(line[i] - 'A') + 27
		//     	}
	    // 		break
	    // 	}
	    // }

    	line1 := fileScanner.Text();
    	fileScanner.Scan()
    	line2 := fileScanner.Text();
    	fileScanner.Scan()
    	line3 := fileScanner.Text();
    	for i := 0; i < len(line1); i++ {
    		if (strings.Index(line2, line1[i:i+1]) >= 0 && strings.Index(line3, line1[i:i+1]) >= 0) {
    			if (line1[i] >= 'a') {
		    		sum += int(line1[i] - 'a') + 1
		    	} else {
		    		sum += int(line1[i] - 'A') + 27
		    	}
		    	break;
    		}
    	}
	}
	fmt.Println(sum);
}




