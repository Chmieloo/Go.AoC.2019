package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day2.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var line string
	var numbers, original []int

	line, err = r.ReadString('\n')
	line = strings.TrimSuffix(line, "\n")
	if err != nil {
		panic(err)
	}
	original, _ = sliceAtoi(strings.Split(line, ","))
	numbers, _ = sliceAtoi(strings.Split(line, ","))

	index := 0

	var p1, p2, p3 int

	// additional condition
	numbers[1] = 12
	numbers[2] = 2

	for numbers[index] != 99 {
		p1 = numbers[index+1]
		p2 = numbers[index+2]
		p3 = numbers[index+3]

		if numbers[index] == 1 {
			numbers[p3] = numbers[p1] + numbers[p2]
		}

		if numbers[index] == 2 {
			numbers[p3] = numbers[p1] * numbers[p2]
		}

		index += 4
	}

	fmt.Println("Answer 1:", numbers[0])

	numbers = original

	pointer := 0
	startPointer := 0

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			numbers, _ = sliceAtoi(strings.Split(line, ","))

			numbers[1] = i
			numbers[2] = j

			pointer = startPointer

			for numbers[pointer] != 99 {
				p1 = numbers[pointer+1]
				p2 = numbers[pointer+2]
				p3 = numbers[pointer+3]

				if numbers[pointer] == 1 {
					numbers[p3] = numbers[p1] + numbers[p2]
				}

				if numbers[pointer] == 2 {
					numbers[p3] = numbers[p1] * numbers[p2]
				}
				pointer += 4
			}
			if numbers[0] == 19690720 {
				fmt.Println("Answer 2:", 100*numbers[1]+numbers[2])
			}
		}
	}
}

func sliceAtoi(sa []string) ([]int, error) {
	si := make([]int, 0, len(sa))
	for _, a := range sa {
		i, err := strconv.Atoi(a)
		if err != nil {
			return si, err
		}
		si = append(si, i)
	}
	return si, nil
}
