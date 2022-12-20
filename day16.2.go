package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	cache := map[string]int{}

	max := 0

	var move func(valve string, time int, open map[string]bool, path []string, val int) (int)
	move = func(valve string, time int, open map[string]bool, path []string, val int) (int) {
		key := ""
		keys := []string{}
		for k, v := range open {
			if (v) {
				keys = append(keys, k)
			}
		}
		sort.Strings(keys)
		for _, k := range keys {
			key = key + ":" + k
		}
		if cache[key] < val {
			cache[key] = val
		}
		if time == 0 {
			return 0;
		}
		ret := 0
		if (!open[valve] && rates[valve]>0) {
			newOpen := map[string]bool{}
			for k, v := range open {
				newOpen[k] = v
			}
			newOpen[valve] = true
			ret = move(valve, time - 1, newOpen, append(path, "*"), val + rates[valve] * (time - 1)) + rates[valve] * (time - 1)
		}
		for _, t := range tunnels[valve] {
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
				ret = Max(ret, move(t, time - 1, open, append(path, t), val))
			}
		}
		return ret
	}

	fmt.Println(move("AA", 26, map[string]bool{}, []string{"AA"}, 0))

	fmt.Println(cache)

	for k1, v1 := range cache {
		for k2, v2 := range cache {
			found := false
			for _, op := range strings.Split(k1, ":") {
				if op!="" && strings.Contains(k2, op) {
					found = true
					break
				}
			}
			if !found {				
				if (max < v1 + v2) {
					max = v1 + v2
					fmt.Println("current max: ", max)
				}
			}
		}
	}

}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}