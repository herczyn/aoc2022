package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func main() {
	readFile, err := os.Open("data8")
	if err != nil {
		fmt.Println(err)
	}

	trees := [][]int{}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		trees = append(trees, []int{})
		for i := 0; i<len(line); i++ {
			v, _ := strconv.Atoi(line[i:i+1])
			trees[len(trees)-1] = append(trees[len(trees)-1], v)
		}
	}
	max :=0
	for x, row := range trees {
		for y, tree := range row {
			if (x>0 && y>0 && x<len(trees)-1 && y<len(row)-1) {
				up:=1
				for (y-up>0 && trees[x][y-up]<tree) {
					up+=1
				}
				down:=1
				for (y+down<len(row)-1 && trees[x][y+down]<tree) {
					down+=1
				}
				left:=1
				for (x-left>0 && trees[x-left][y]<tree) {
					left+=1
				}
				right:=1
				for (x+right<len(trees)-1 && trees[x+right][y]<tree) {
					right+=1
				}
				fmt.Println(x, y, up*down*left*right)
				if (up*down*left*right) > max {
					max = up*down*left*right
				}
			}
		}
	}
	fmt.Println(max)

	// tree_no := 0
	// for x, row := range trees {
	// 	for y, tree := range row {
	// 		found := false
	// 		for i:=0; i<x; i++ {
	// 			if (trees[i][y]>=tree) {
	// 				found = true
	// 			}
	// 		}
	// 		if (!found) {
	// 			tree_no += 1
	// 		} else {
	// 			found = false
	// 			for i:=0; i<y; i++ {
	// 				if (trees[x][i]>=tree) {
	// 					found = true
	// 				}
	// 			}
	// 			if (!found) {
	// 				tree_no += 1
	// 			} else {
	// 				found = false
	// 				for i:=x+1; i<len(trees); i++ {
	// 					if (trees[i][y]>=tree) {
	// 						found = true
	// 					}
	// 				}
	// 				if (!found) {
	// 					tree_no += 1
	// 				} else {
	// 					found = false
	// 					for i:=y+1; i<len(row); i++ {
	// 						if (trees[x][i]>=tree) {
	// 							found = true
	// 						}
	// 					}
	// 					if (!found) {
	// 						tree_no += 1
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// fmt.Println(tree_no)
}

