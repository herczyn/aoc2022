package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("data16")
	if err != nil {
		fmt.Println(err)
	}

	rates := map[string]int{}
	tunnels := map[string][]string{}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)
  
	for fileScanner.Scan() {
		line := fileScanner.Text();
		valve:=strings.Split(line, " ")[1]
		rate, _ :=strconv.Atoi(strings.Split(strings.Split(line, "=")[1], ";")[0])
		var t []string
		if (strings.Contains(line, "to valves")) {
			t = strings.Split(strings.Split(line, "to valves ")[1], ", ")
		} else {
			t = []string{strings.Split(line, "to valve ")[1]}
		}

		rates[valve] = rate
		tunnels[valve] = t

		fmt.Println(valve, rate, t)
	}

	var move func(valve string, time int, open map[string]bool, path []string) (int)
	move = func(valve string, time int, open map[string]bool, path []string) (int) {
		if time == 0 {
			return 0;
		}
		ret := 0
		if (!open[valve] && rates[valve]>0) {
			// fmt.Println("path", path, "time left", time, "opening", valve)
			newOpen := map[string]bool{}
			for k, v := range open {
				newOpen[k] = v
			}
			newOpen[valve] = true
			ret = move(valve, time - 1, newOpen, append(path, "*")) + rates[valve] * (time - 1)
		}
		for _, t := range tunnels[valve] {
			// fmt.Println("path", path, "time left", time, "going to valve", t)
			found := false
			for i := len(path)-1; i>=0; i-- {
				if path[i] == "*" {
					break;
				}
				if path[i] == t {
					found = true
					break;
				}
			}
			if !found {
				ret = Max(ret, move(t, time - 1, open, append(path, t)))
			}
		}
		return ret
	}

	fmt.Println(move("AA", 30, map[string]bool{}, []string{"AA"}))
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}