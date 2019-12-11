package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
	"strconv"
)

func main() {
	f, err := os.Open("day8.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var line string

	line, err = r.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	if err != nil {
		panic(err)
	}

	runes := []rune(line)

	w := 25
	h := 6

	len := w * h

	var image [6]string

	startingIndexWithLeastZeros := 0
	zeros := len
	var slice string
	var resultSlice string

	for i := 0; i < utf8.RuneCountInString(line); i += len {
		slice = string(runes[i : i+len])
		if resultSlice == "" {
			// copy the original slice
			resultSlice = slice
		} else {
			// Compare slices
			for pos, char := range slice {
				orig, _ := strconv.Atoi(string(resultSlice[pos]))
				new, _ := strconv.Atoi(string(char))
				if new < orig && orig == 2 {
					resultSlice = resultSlice[:pos] + strconv.Itoa(new) + resultSlice[pos+1:]
				}
			}
		}
		
		numZeros := strings.Count(slice, "0")
		if numZeros < zeros {
			startingIndexWithLeastZeros = i
			zeros = numZeros
		}


		row := 0
		for j := 0; j < len; j += w {
			//fmt.Println(string(runes[j : j+w]))
			image[row] = string(resultSlice[j : j+w])
			row++
		}
	}

	s := string(runes[startingIndexWithLeastZeros : startingIndexWithLeastZeros+len])
	numOnes := strings.Count(s, "1")
	numTwos := strings.Count(s, "2")

	fmt.Println("Answer part 1:", numOnes*numTwos)
	fmt.Println("Answer part 2:")
	for _, r := range image {
		l := strings.Replace(r, "0", " ", -1)
		l = strings.Replace(l, "1", "#", -1)
		fmt.Println(l)
	}
}
