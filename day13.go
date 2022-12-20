package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	readFile, err := os.Open("data13")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
 
	fileScanner.Split(bufio.ScanLines)

	getInt := func(a string) (int, int) {
		if (a[0] == '[') {
			return -1, -1
		} else {
			i := 0
			for ; i<len(a) && a[i] != ','; i++ {}
			ret, _ := strconv.Atoi(a[0:i])
			return ret, i
		}
	}

	getArray := func(a string) (string, int) {
		pars := 1
		i := 1
		for ; pars>0; i++ {
			if a[i] == '[' {
				pars++
			}
			if a[i] == ']' {
				pars--
			}
		} 
		return a[1:i-1], i
	}
	var compare func(a string, b string) (int)
	compare = func(a string, b string) (int) {
		if (len(a)==0 && len(b)==0) {
			return 0
		} else if len(a) == 0 {
			return 1
		} else if len(b) == 0 {
			return -1
		}
		n, s := getInt(a)
		m, z := getInt(b)
		var q,r string
		if n>=0 && m>=0 {
			if n<m {
				return 1
			} else if n>m {
				return -1
			}
		} else if n>=0 {
			r, z = getArray(b)
			c := compare(strconv.Itoa(n), r)
			if (c != 0) {
				return c
			}
		} else if m>=0 {
			q, s = getArray(a)
			c := compare(q, strconv.Itoa(m))
			if (c != 0) {
				return c
			}
		} else {
			q, s = getArray(a)
			r, z = getArray(b)
			c := compare(q, r)
			if (c != 0) {
				return c
			}
		}
		if s==len(a) && s==len(b) {
			return 0;
		} else if s==len(a) {
			return 1
		} else if z==len(b) {
			return -1
		} else {
			return compare(a[s+1:], b[z+1:])
		}
	}
	// i:=0
	// ret:=0
	// for fileScanner.Scan() {
	// 	i++
	// 	line1 := fileScanner.Text();
	// 	fileScanner.Scan()
	// 	line2 := fileScanner.Text();
	// 	fileScanner.Scan()
	// 	c:=compare(line1, line2)
	// 	fmt.Println(line1, line2, c)
	// 	if c==1 {
	// 		ret+=i
	// 	}
	// }
	// fmt.Println(ret)
	lines := []string{"[[2]]", "[[6]]"}
	for fileScanner.Scan() {
		line := fileScanner.Text();
		if line != "" {
			lines = append(lines, line)
		}
	}
	found := true
	for found {
		found = false
		for i := 0; i < len(lines)-1; i++ {
			if compare(lines[i], lines[i+1])==-1 {
				t := lines[i]
				lines[i] = lines[i+1]
				lines[i+1] = t
				found = true
			}
		}
		fmt.Println(lines)
	}
	i2 := 0
	i6 := 0
	for i, line := range lines {
		if line=="[[2]]" {
			i2=i
		}
		if line=="[[6]]" {
			i6=i
		}
	}
	fmt.Println((i2+1)*(i6+1))
}
