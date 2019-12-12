package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	start := 387638
	end := 919123

	validNumbers := 0

	for i := start; i <= end; i++ {
		if checkNumber(i) {
			validNumbers++
		}
	}

	fmt.Println("Answer 1:", validNumbers)

	validNumbers = 0
	for i := start; i <= end; i++ {
		if checkNumber2(i) {
			validNumbers++
		}
	}

	fmt.Println("Answer 2:", validNumbers)
}

func checkNumber(i int) bool {
	s := strconv.Itoa(i)

	hasDouble := false
	ordered := true
	for j := 1; j < len(s); j++ {
		if s[j] == s[j-1] {
			hasDouble = true
		}
		sn, _ := strconv.Atoi(string(s[j]))
		sp, _ := strconv.Atoi(string(s[j-1]))
		if sn < sp {
			ordered = false
			break
		}
	}

	if hasDouble && ordered {
		return true
	} else {
		return false
	}
}

func checkNumber2(i int) bool {
	s := strconv.Itoa(i)

	for j := 1; j < len(s); j++ {
		sn, _ := strconv.Atoi(string(s[j]))
		sp, _ := strconv.Atoi(string(s[j-1]))
		if sn < sp {
			return false
		}
	}

	var c map[int]int
	c = map[int]int{}
	doubles := 0
	nums := strings.Split(s, "")
	for j := 0; j < len(nums); j++ {
		str, _ := strconv.Atoi(nums[j])
		if c[str] == 0 {
			c[str]++
		} else {
			c[str]++
			doubles++
		}
	}

	if doubles >= 1 {
		for _, val := range c {
			if val == 2 {
				//fmt.Println(s, c, key, val)
				return true
			}
		}
	}

	return false
}
