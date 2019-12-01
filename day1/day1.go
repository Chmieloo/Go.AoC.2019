package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day1.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	r := bufio.NewReader(f)

	var sum, sum2 int
	var mass int
	var requirement int
	var subsum int

	var line string
	for {
		subsum = 0
		line, err = r.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		mass, _ = strconv.Atoi(line)
		requirement = int(math.Floor(float64(mass / 3)) - 2)
		sum += int(math.Floor(float64(mass / 3)) - 2)

		for requirement > 0 {
			subsum += requirement
			requirement = int(math.Floor(float64(requirement / 3)) - 2)
		}

		sum2 += subsum
	}

	fmt.Println("Answer 1:", sum)
	fmt.Println("Answer 2:", sum2)
}